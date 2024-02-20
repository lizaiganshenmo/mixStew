package service

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/article/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/article/rpc"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"golang.org/x/sync/errgroup"
)

func (s *ArticleService) GetArticleComment(req *article.GetArticleCommentReq) ([]*article.Comment, error) {
	comments, err := rpc.GetArticleComment(s.ctx, &interaction.GetCommentReq{
		TargertArticleId: req.ArticleId,
	})
	// 每个评论查对应作者信息
	commentList := make([]*article.Comment, len(comments))
	eg, ctx := errgroup.WithContext(s.ctx)
	for i, v := range comments {
		idx := i
		t := v
		eg.Go(
			func() error {
				defer func() {
					if r := recover(); r != nil {
						klog.CtxWarnf(ctx, "recover panic.err:%+v", err)
					}
				}()

				userInfo, err := rpc.GetUser(ctx, &user.GetUserReq{
					Uid: t.Uid,
				})
				if err != nil {
					klog.CtxWarnf(s.ctx, "rpc.GetUser fail,errr:%+v", err)
				}

				var following bool
				following, err = rpc.IsFollow(s.ctx, &follow.FollowReq{
					Uid:       t.Uid,
					FollowUid: t.TargertUid,
				})
				if err != nil {
					klog.CtxWarnf(s.ctx, "rpc.IsFollow fail,errr:%+v", err)
				}

				commentList[idx] = pack.Comment(t, userInfo, following)

				return nil

			},
		)
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return commentList, err

}
