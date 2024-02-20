package service

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/cache"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
	"github.com/lizaiganshenmo/mixStew/cmd/article/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/article/rpc"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/utils"
	"golang.org/x/sync/errgroup"
)

func (s *ArticleService) MGetArticle(req *article.MGetArticleReq) (articleList []*article.Article, err error) {
	// s1 根据是否有文章tag查 article_ids
	var articleIds []int64
	if req.Tag != nil {
		var ids []int64
		ids, err = cache.GetArticleIdsByTag(s.ctx, *req.Tag)
		if err != nil {
			klog.CtxWarnf(s.ctx, "cache.GetArticleIdsByTag fail. err:%+v", err)
			return
		}

		if len(ids) == 0 {
			ids, err = db.GetArticleIdsByTag(s.ctx, *req.Tag)
			if err != nil {
				klog.CtxWarnf(s.ctx, " db.GetArticleIdsByTag fail. err:%+v", err)
				return
			}

			cache.SetArticleIdsByTag(s.ctx, *req.Tag, ids)
		}

		articleIds = ids

	}

	// s2 根据是否查询喜爱文章id
	var favIds []int64
	if req.Favorited != nil && *req.Favorited {
		var ids []int64
		ids, err = s.FavoriteList(req.Uid)
		if err != nil {
			klog.CtxWarnf(s.ctx, "s.FavoriteList fail. err:%+v", err)
			return
		}

		favIds = ids
	}

	if len(articleIds) != 0 && len(favIds) != 0 {
		articleIds = utils.GetCommonElems(articleIds, favIds)
	} else if len(articleIds) == 0 && len(favIds) != 0 {
		articleIds = favIds
	}

	// s3 查询文章
	var aInfos []*db.ArticleInfo
	if len(articleIds) != 0 {
		aInfos, err = db.QueryArticle(s.ctx, db.WithArticleIds(articleIds),
			db.WithLimit(int(req.Limit)), db.WithOffest(int(req.Offest)), db.WithOrder("created_at desc"))
	} else {
		aInfos, err = db.QueryArticle(s.ctx, db.WithLimit(int(req.Limit)), db.WithOffest(int(req.Offest)), db.WithOrder("created_at desc"))
	}
	klog.Warnf("len aInfos is %d\n", len(aInfos))

	// s4 查询文章的tag  、点赞数、 作者信息
	articleList = make([]*article.Article, len(aInfos))
	eg, ctx := errgroup.WithContext(s.ctx)
	for i, v := range aInfos {
		idx := i
		tmpInfo := v
		articleId := tmpInfo.ArticleId
		authorUid := tmpInfo.Uid
		eg.Go(func() error {
			defer func() {
				if r := recover(); r != nil {
					klog.CtxWarnf(ctx, "recover panic.err:%+v", err)
				}
			}()

			// 查询文章tag
			var tagList []string
			var err error
			tagList, err = cache.GetArticleTags(ctx, articleId)
			if err != nil {
				klog.CtxWarnf(ctx, "cache.GetArticleTags fail. err:%+v", err)
			}
			if len(tagList) == 0 {
				tagList, err = db.GetTagsByArticleId(ctx, articleId)
				if err != nil {
					klog.CtxWarnf(ctx, "db.GetTagsByArticleId fail. err:%+v", err)
				} else if len(tagList) != 0 {
					err = cache.SetArticleTags(ctx, articleId, tagList)
					if err != nil {
						klog.CtxWarnf(ctx, "cache.SetArticleTags fail. err:%+v", err)
					}
				}

			}

			// 查询点赞数
			var favCount int64
			favCount, err = cache.GetArticleFavoriteCount(ctx, articleId)
			if err != nil {
				favCount, err = db.GetFavoriteCountByArticleId(ctx, articleId)
				if err != nil {
					klog.CtxWarnf(ctx, "db.GetFavoriteCountByArticleId fail. err:%+v", err)
				} else {
					err = cache.SetArticleFavouriteCount(ctx, articleId, favCount)
					if err != nil {
						klog.CtxWarnf(ctx, "cache.SetArticleFavouriteCount fail. err:%+v", err)
					}
				}
			}
			klog.CtxWarnf(ctx, "favCount:%d", favCount)

			// 查询作者信息
			var userInfo *user.User
			userInfo, err = rpc.GetUser(ctx, &user.GetUserReq{Uid: authorUid})
			if err != nil {
				klog.CtxWarnf(ctx, "rpc.GetUser fail. err:%+v", err)
			}

			// 查询是否关注
			var following bool
			following, err = rpc.IsFollow(ctx, &follow.FollowReq{Uid: req.Uid, FollowUid: authorUid})
			if err != nil {
				klog.CtxWarnf(ctx, "rpc.IsFollow fail. err:%+v", err)
			}

			// 组装信息
			articleList[idx] = pack.Article(*tmpInfo, userInfo, following, *req.Favorited, favCount, tagList)

			return nil
		})

	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return

}

func (s *ArticleService) MGetFeedArticle(req *article.MGetFeedArticleReq) (articleList []*article.Article, err error) {
	// s1 查询关注uids
	var UIds []int64
	UIds, err = rpc.FollowList(s.ctx, &follow.FollowListReq{Uid: req.Uid})
	if err != nil {
		return
	}

	if len(UIds) == 0 {
		return
	}

	// s2 根据 uids 查询文章
	var aInfos []*db.ArticleInfo
	aInfos, err = db.QueryArticle(s.ctx, db.WithUids(UIds),
		db.WithLimit(int(req.Limit)), db.WithOffest(int(req.Offest)), db.WithOrder("created_at desc"))

	// s3 查询喜爱文章列表
	articleFavIds, err1 := s.FavoriteList(req.Uid)
	if err1 != nil {
		klog.CtxWarnf(s.ctx, "s.FavoriteList fail. err:%+v", err1)
	}
	articleFavIdsMap := make(map[int64]struct{}, len(articleFavIds))
	for _, v := range articleFavIds {
		articleFavIdsMap[v] = struct{}{}
	}

	// s4 查询文章的tag  、点赞数、 作者信息
	articleList = make([]*article.Article, len(aInfos))
	eg, ctx := errgroup.WithContext(s.ctx)
	for i, v := range aInfos {
		idx := i
		tmpInfo := v
		articleId := tmpInfo.ArticleId
		authorUid := tmpInfo.Uid
		eg.Go(func() error {
			defer func() {
				if r := recover(); r != nil {
					klog.CtxWarnf(ctx, "recover panic.err:%+v", err)
				}
			}()

			// 查询文章tag
			var tagList []string
			var err error
			tagList, err = cache.GetArticleTags(ctx, articleId)
			if err != nil {
				klog.CtxWarnf(ctx, "cache.GetArticleTags fail. err:%+v", err)
			}
			if len(tagList) == 0 {
				tagList, err = db.GetTagsByArticleId(ctx, articleId)
				if err != nil {
					klog.CtxWarnf(ctx, "db.GetTagsByArticleId fail. err:%+v", err)
				} else if len(tagList) != 0 {
					err = cache.SetArticleTags(ctx, articleId, tagList)
					if err != nil {
						klog.CtxWarnf(ctx, "cache.SetArticleTags fail. err:%+v", err)
					}
				}

			}

			// 查询点赞数
			var favCount int64
			favCount, err = cache.GetArticleFavoriteCount(ctx, articleId)
			if err != nil {
				favCount, err = db.GetFavoriteCountByArticleId(ctx, articleId)
				if err != nil {
					klog.CtxWarnf(ctx, "db.GetFavoriteCountByArticleId fail. err:%+v", err)
				} else {
					err = cache.SetArticleFavouriteCount(ctx, articleId, favCount)
					if err != nil {
						klog.CtxWarnf(ctx, "cache.SetArticleFavouriteCount fail. err:%+v", err)
					}
				}
			}

			// 查询作者信息
			var userInfo *user.User
			userInfo, err = rpc.GetUser(ctx, &user.GetUserReq{Uid: authorUid})
			if err != nil {
				klog.CtxWarnf(ctx, "rpc.GetUser fail. err:%+v", err)
			}

			// 查询是否关注
			var following bool
			following, err = rpc.IsFollow(ctx, &follow.FollowReq{Uid: req.Uid, FollowUid: authorUid})
			if err != nil {
				klog.CtxWarnf(ctx, "rpc.IsFollow fail. err:%+v", err)
			}

			// 是否点赞该文章
			var favorited bool
			if _, ok := articleFavIdsMap[articleId]; ok {
				favorited = true
			}

			// 组装信息
			articleList[idx] = pack.Article(*tmpInfo, userInfo, following, favorited, favCount, tagList)

			return nil
		})

	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return

}
