package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	validator "github.com/go-playground/validator/v10"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

var (
	vld = validator.New()
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

// valdate req is valid
func ValidateFunc(c *app.RequestContext, req interface{}) {
	err := vld.Struct(req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
	}
}
