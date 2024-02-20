package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/service"
	interaction "github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// CreateComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CreateComment(ctx context.Context, req *interaction.CreateCommentReq) (resp *interaction.CreateCommentResp, err error) {
	resp = new(interaction.CreateCommentResp)

	err = service.NewInteractionService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "inetraction.CreateComment fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// DeleteComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) DeleteComment(ctx context.Context, req *interaction.DeleteCommentReq) (resp *interaction.DeleteCommentResp, err error) {
	resp = new(interaction.DeleteCommentResp)

	err = service.NewInteractionService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "inetraction.DeleteComment fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) GetComment(ctx context.Context, req *interaction.GetCommentReq) (resp *interaction.GetCommentResp, err error) {
	resp = new(interaction.GetCommentResp)
	commentList, err := service.NewInteractionService(ctx).GetComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "inetraction.GetComment fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.Comments = pack.Comments(commentList)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// CreateCommentReply implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CreateCommentReply(ctx context.Context, req *interaction.CreateCommentReplyReq) (resp *interaction.CreateCommentReplyResp, err error) {
	resp = new(interaction.CreateCommentReplyResp)

	err = service.NewInteractionService(ctx).CreateCommentReply(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "inetraction.CreateCommentReply fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// DeleteCommentReply implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) DeleteCommentReply(ctx context.Context, req *interaction.DeleteCommentReplyReq) (resp *interaction.DeleteCommentReplyResp, err error) {
	resp = new(interaction.DeleteCommentReplyResp)

	err = service.NewInteractionService(ctx).DeleteCommentReply(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "inetraction.DeleteCommentReply fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetCommentReply implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) GetCommentReply(ctx context.Context, req *interaction.GetCommentReplyReq) (resp *interaction.GetCommentReplyResp, err error) {
	resp = new(interaction.GetCommentReplyResp)
	commentList, err := service.NewInteractionService(ctx).GetCommentReply(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "inetraction.GetComment fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.CommentReplys = pack.Comments(commentList)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
