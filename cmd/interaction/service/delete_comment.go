package service

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/mq"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

func (s *InteractionService) DeleteComment(req *interaction.DeleteCommentReq) error {
	// s1 判断是否是本人评论
	comments, err := db.GetCommentsByCommentId(s.ctx, req.CommentId)
	if err != nil {
		return err
	}

	if len(comments) == 0 || comments[0].Uid != req.Uid {
		return errno.CommentNotExistErr
	}

	// s2 推送至mq
	cmt := mq.CommentMiddle{
		Comment: &db.CommentInfo{
			Uid:       req.Uid,
			CommentId: req.CommentId,
		},
		Operation: mq.DeleteOperation,
	}
	c, _ := sonic.MarshalString(&cmt)

	err = mq.PublishComment(s.ctx, c)
	if err != nil {
		klog.CtxWarnf(s.ctx, "mq.PublishComment fail. err: %+v", err)
		return err
	}

	return nil
}
