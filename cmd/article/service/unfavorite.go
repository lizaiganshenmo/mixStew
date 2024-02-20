package service

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/cache"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

func (s *ArticleService) UnFavourite(req *article.UnFavoriteArticleReq) error {
	// s1 是否已经不喜欢
	exist, err := cache.IsExistArticle(s.ctx, req.ArticleId)
	if err != nil {
		klog.CtxWarnf(s.ctx, "cache.IsExistArticle fail. err: %+v req: %+v", err, req)
		return err
	}
	// key 存在
	if exist == 1 {
		hasLike, err := cache.IsFavorited(s.ctx, req.ArticleId, req.Uid)
		if err != nil {
			klog.CtxWarnf(s.ctx, "cache.IsFavorited fail. err: %+v req: %+v", err, req)
			return err
		}

		if !hasLike {
			return errno.ArticleFavoriteAlreadyNotExistErr
		}

	}

	// // s2 reduce favorite info to cache
	// if err := cache.ReduceFavorite(s.ctx, req.ArticleId, req.Uid); err != nil {
	// 	klog.CtxWarnf(s.ctx, "cache.ReduceFavorite fail. err: %+v req: %+v", err, req)
	// 	return err
	// }

	// s2 update like status to db
	if err := db.UnFavorite(s.ctx, req.ArticleId, req.Uid); err != nil {
		klog.CtxWarnf(s.ctx, "db.UnFavorite fail. err: %+v req: %+v", err, req)
		return err
	}

	// todo : del cache task -> mq to consume (with retry)
	// s3 del cache info
	if err := cache.DelArticleFavoriteInfo(s.ctx, req.ArticleId, req.Uid); err != nil {
		klog.CtxWarnf(s.ctx, "cache.DelArticleFavoriteInfo fail. err: %+v req: %+v", err, req)
		return err
	}

	return nil
}
