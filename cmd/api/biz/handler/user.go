package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/rpc"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/types/req"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var req req.CreateUserReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	ValidateFunc(c, req)

	err := rpc.CreateUser(ctx, &user.CreateUserReq{
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
