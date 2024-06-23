package web

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserRegisterRequest struct {
	Name        string `validate:"required" json:"name"`
	Username    string `validate:"required" json:"username"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required" json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type AdminRegisterRequest struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
