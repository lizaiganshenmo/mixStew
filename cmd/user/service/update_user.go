package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/user/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/utils"
)

func (s *UserService) UpdateUser(req *user.UpdateUserReq) error {
	err := db.UpdateUser(s.ctx, &db.User{
		UserName: req.Username,
		Email:    req.Email,
		Bio:      req.Bio,
		Password: utils.EncryptPassword(req.Password),
		Image:    req.Image,
	})
	if err != nil {
		return err
	}

	return nil
}
