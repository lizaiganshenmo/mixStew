package db

import (
	"context"

	"gorm.io/gorm"
)

const (
	UserTableName = "user"
)

type User struct {
	gorm.Model
	Uid      int64  `json:"uid" gorm:"column:uid"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Bio      string `json:"bio" gorm:"column:bio"`
	Image    string `json:"image" gorm:"column:image"`
}

func (u *User) TableName() string {
	return UserTableName
}

// create user info
func CreateUser(ctx context.Context, users []*User) error {
	return MySQLMixStewDB.WithContext(ctx).Create(users).Error
}

// query user info by email
func QueryUser(ctx context.Context, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := MySQLMixStewDB.WithContext(ctx).Where("email = ?", email).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
