package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
)

func (s *InteractionService) GetComment(req *interaction.GetCommentReq) (commentList []*db.CommentInfo, err error) {
	return db.GetCommentsByArticleId(s.ctx, req.TargertArticleId)
}
