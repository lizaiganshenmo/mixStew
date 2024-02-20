package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/article/rpc"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
)

func (s *ArticleService) CommentArticle(req *article.CommentArticleReq) error {
	return rpc.CommentArticle(s.ctx, &interaction.CreateCommentReq{
		Uid:              req.Uid,
		TargertArticleId: req.ArticleId,
		TargertUid:       req.TargetUid,
		CommentBody:      req.Body,
	})
}
