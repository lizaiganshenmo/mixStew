package pack

import (
	"github.com/lizaiganshenmo/mixStew/cmd/user/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		UserRegister: &user.CreateUserReq{
			UserAuth: &user.UserAuth{Email: u.Email},
			Username: u.UserName,
		},
		Bio:   u.Bio,
		Image: u.Image,
	}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
