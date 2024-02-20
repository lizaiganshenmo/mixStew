package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	UserNotExistErrCode        = 10005
	UserNotLoginErrCode        = 10006

	// follow
	FollowYourselfErrorCode = 11001

	// article
	ArticleFavoriteAlreadyExist    = 12001
	ArticleFavoriteAlreadyNotExist = 12001

	// interaction
	SensitiveWordsErrCode  = 13001
	CommentNotExistErrCode = 13002
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, "user not exit")
	UserNotLoginErr        = NewErrNo(UserNotLoginErrCode, "user not login")

	// follow
	FollowYourselfError = NewErrNo(FollowYourselfErrorCode, "can not follow yourself")

	// article
	ArticleFavoriteAlreadyExistErr    = NewErrNo(ArticleFavoriteAlreadyExist, "already favority article")
	ArticleFavoriteAlreadyNotExistErr = NewErrNo(ArticleFavoriteAlreadyNotExist, "already unfavority article")

	// interaction
	SensitiveWordsErr  = NewErrNo(SensitiveWordsErrCode, "existed sensitive words")
	CommentNotExistErr = NewErrNo(CommentNotExistErrCode, "comment not existed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
