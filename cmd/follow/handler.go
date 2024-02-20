package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/follow/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/follow/service"
	follow "github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Follow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Follow(ctx context.Context, req *follow.FollowReq) (resp *follow.FollowResp, err error) {
	resp = new(follow.FollowResp)
	if req.Uid == 0 || req.FollowUid == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.NewFollowService(ctx).Follow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "follow.Follow fail. err : %+v, req: %+v", err, req)
		return
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// UnFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) UnFollow(ctx context.Context, req *follow.FollowReq) (resp *follow.FollowResp, err error) {
	resp = new(follow.FollowResp)
	if req.Uid == 0 || req.FollowUid == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.NewFollowService(ctx).Unfollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "follow.UnFollow fail. err : %+v, req: %+v", err, req)
		return
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// IsFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) IsFollow(ctx context.Context, req *follow.FollowReq) (resp *follow.IsFollorResp, err error) {
	resp = new(follow.IsFollorResp)
	// if req.Uid == 0 || req.FollowUid == 0 {
	// 	resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	// 	return
	// }

	isFollow, err := service.NewFollowService(ctx).IsFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "follow.IsFollow fail. err : %+v, req: %+v", err, req)
		return
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFollowing = isFollow
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListReq) (resp *follow.FollowListResp, err error) {
	resp = new(follow.FollowListResp)
	if req.Uid == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uids, err := service.NewFollowService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "follow.FollowList fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Uids = uids
	return

}
