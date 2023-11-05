package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/user/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/user/service"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CheckUser implements user.UserService.
func (*UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserReq) (resp *user.CheckUserResp, err error) {
	resp = new(user.CheckUserResp)

	uid, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		klog.CtxWarnf(ctx, "user.CheckUser fail. err : %+v, req: %+v", err, req)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Uid = uid
	return resp, nil
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.CreateUserResp, err error) {
	resp = new(user.CreateUserResp)

	if len(req.Username) == 0 || len(req.UserAuth.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "user.CreateUser fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	resp = new(user.GetUserResp)

	if req.Uid == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewUserService(ctx).GetUser(req)
	fmt.Printf("ctx kv is: %s: %v\n", constants.RequestIdKey, ctx.Value(constants.RequestIdKey))

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "user.GetUser fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = pack.User(user)
	return resp, nil
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	resp = new(user.UpdateUserResp)
	if len(req.Email) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUserService(ctx).UpdateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "user.UpdateUser fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
