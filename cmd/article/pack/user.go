package pack

import (
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
)

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

func User(data *user.User) *UserInfo {
	return &UserInfo{
		Email:    data.UserRegister.UserAuth.Email,
		Password: data.UserRegister.UserAuth.Password,
		Username: data.UserRegister.Username,
		Bio:      data.Bio,
		Image:    data.Image,
	}

}
