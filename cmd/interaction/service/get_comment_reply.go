package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
)

func (s *InteractionService) GetCommentReply(req *interaction.GetCommentReplyReq) (commentList []*db.CommentInfo, err error) {
	return db.GetCommentsByCommentId(s.ctx, req.TargertCommentId)
}
