package pack

import (
	"errors"
	"time"

	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *interaction.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *interaction.BaseResp {
	return &interaction.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
