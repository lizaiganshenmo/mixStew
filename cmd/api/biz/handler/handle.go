package handler

import (
	"regexp"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	validator "github.com/go-playground/validator/v10"
	"github.com/hertz-contrib/requestid"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

var (
	vld             = validator.New()
	valRegisterOnce sync.Once
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	LogId   string      `json:"log_id"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		LogId:   requestid.Get(c),
		Data:    data,
	})
}

// valdate req is valid
func ValidateFunc(c *app.RequestContext, req interface{}) error {
	valRegisterOnce.Do(func() {
		vld.RegisterValidation("regexEmail", validateEmail)
	})
	err := vld.Struct(req)
	if err != nil {
		// SendResponse(c, errno.ParamErr, nil)
		return err
	}

	return nil
}

// validate email is valid
func validateEmail(fl validator.FieldLevel) bool {
	pattern := `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	match, _ := regexp.MatchString(pattern, fl.Field().String())
	return match
}
