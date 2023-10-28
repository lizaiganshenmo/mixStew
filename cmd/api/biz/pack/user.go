package pack

import (
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"

	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/types/resp"
)

func User(data *user.User) *resp.User {
	return &resp.User{
		Email:    data.UserRegister.UserAuth.Email,
		Password: data.UserRegister.UserAuth.Password,
		Username: data.UserRegister.Username,
		Bio:      data.Bio,
		Image:    data.Image,
	}

}
