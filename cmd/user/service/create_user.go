package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/user/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/lizaiganshenmo/mixStew/library/utils"
)

func (s *UserService) CreateUser(req *user.CreateUserReq) error {
	// s1 查询该email是否已注册
	users, err := db.QueryUser(s.ctx, req.UserAuth.Email)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return errno.UserAlreadyExistErr
	}
	// s2 密码加密
	password := utils.EncryptPassword(req.UserAuth.Password)
	if password == "" {
		return errno.ParamErr
	}
	// s3 创建用户账号
	return db.CreateUser(s.ctx, []*db.User{{
		Uid:      db.GenUID(),
		UserName: req.Username,
		Password: password,
		Email:    req.UserAuth.Email,
		// Bio:      "",
		// Image:    "",
	},
	})

}
