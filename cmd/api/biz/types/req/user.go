package req

type UserParamReq struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		PassWord string `json:"password" validate:"required"`
	} `json:"user" validate:"required"`
}

type CreateUserReq struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		PassWord string `json:"password" validate:"required"`
		UserName string `json:"username" validate:"required"`
	} `json:"user" validate:"required"`
}
