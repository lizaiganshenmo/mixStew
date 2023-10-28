package req

type UserParamReq struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		PassWord string `json:"password" validate:"required"`
	} `json:"user" validate:"required"`
}

type CreateUserReq struct {
	User struct {
		Email    string `json:"email" validate:"required,regexEmail"`
		PassWord string `json:"password" validate:"required,gte=6,lte=16"`
		UserName string `json:"username" validate:"required"`
	} `json:"user" validate:"required"`
}

type UpdateUserReq struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		PassWord string `json:"password" validate:"required,gte=6,lte=16"`
		UserName string `json:"username" validate:"required"`
		Bio      string `json:"bio" validate:"required"`
		Image    string `json:"image" validate:"required"`
	} `json:"user" validate:"required"`
}
