package service

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/mq"
	sensitiveword "github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/sensitive_word"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

func (s *InteractionService) CreateCommentReply(req *interaction.CreateCommentReplyReq) error {
	// s1 评论 敏感词库初筛
	if sensitiveword.HasSensitiveWord(req.CommentBody) {
		return errno.SensitiveWordsErr
	}

	// s2 推送至mq
	cmt := mq.CommentMiddle{
		Comment: &db.CommentInfo{
			Uid:             req.Uid,
			CommentId:       db.GenCommentId(),
			TargetArticleId: req.TargertArticleId,
			TargetUid:       req.TargertUid,
			TargetCommentId: req.TargertCommentId,
			CommentType:     db.CommentTypeTwo,
			Status:          db.PassStatus,
			Body:            req.CommentBody,
		},
		Operation: mq.CreateOperation,
	}
	c, _ := sonic.MarshalString(&cmt)

	err := mq.PublishComment(s.ctx, c)
	if err != nil {
		klog.CtxWarnf(s.ctx, "mq.PublishComment fail. err: %+v", err)
		return err
	}

	return nil
}
