package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/rpc"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/types/req"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/lizaiganshenmo/mixStew/library/utils"
)

func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var req req.CreateUserReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.CreateUser(ctx, &user.CreateUserReq{
		UserAuth: &user.UserAuth{
			Email:    req.User.Email,
			Password: req.User.PassWord,
		},
		Username: req.User.UserName,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.CreateUser err: %+v", err)
		return
	}

	SendResponse(c, errno.Success, nil)
}

func GetUser(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	user, err := rpc.GetUser(ctx, &user.GetUserReq{Uid: uid})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.GetUser err: %+v", err)
		return
	}

	SendResponse(c, errno.Success, pack.User(user))
}

func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var req req.UpdateUserReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.UpdateUser(ctx, &user.UpdateUserReq{
		Email:    req.User.Email,
		Username: req.User.UserName,
		Password: req.User.PassWord,
		Bio:      req.User.Bio,
		Image:    req.User.Image,
	})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.UpdateUser err: %+v", err)
		return
	}

	SendResponse(c, errno.Success, nil)

}
