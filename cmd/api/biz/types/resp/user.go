package resp

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
