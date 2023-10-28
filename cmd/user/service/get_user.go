package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/user/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

func (s *UserService) GetUser(req *user.GetUserReq) (*db.User, error) {
	users, err := db.QueryUserByUid(s.ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}

	return users[0], nil
}
