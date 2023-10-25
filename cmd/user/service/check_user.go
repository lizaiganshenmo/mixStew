package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/user/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/lizaiganshenmo/mixStew/library/utils"
)

func (s *UserService) CheckUser(req *user.CheckUserReq) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.User_Auth.Email)
	if err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}

	u := users[0]
	if u.Password != utils.EncryptPassword(req.User_Auth.Password) {
		return 0, errno.AuthorizationFailedErr
	}

	return u.Uid, nil
}
