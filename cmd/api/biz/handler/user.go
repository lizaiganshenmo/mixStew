package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/rpc"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/types/req"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/errno"
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
		hlog.Warnf("rpc.CreateUser err: %+v", err)
		return
	}

	SendResponse(c, errno.Success, nil)
}

func GetUser(ctx context.Context, c *app.RequestContext) {
	var uid int64
	t, ok := c.Get(constants.IdentityKey)
	hlog.Warnf("start get constants.IdentityKey: %+v  ok: %t", t, ok)
	if !ok {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	uid = int64(t.(float64))
	user, err := rpc.GetUser(ctx, &user.GetUserReq{Uid: uid})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.Warnf("rpc.GetUser err: %+v", err)
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
		hlog.Warnf("rpc.UpdateUser err: %+v", err)
		return
	}

	SendResponse(c, errno.Success, nil)

}
