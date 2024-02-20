package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/article/rpc"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
)

func (s *ArticleService) DeleteComment(req *article.DeleteCommentReq) error {
	return rpc.DeleteComment(s.ctx, &interaction.DeleteCommentReq{
		Uid:       req.Uid,
		CommentId: req.CommentId,
	})
}
