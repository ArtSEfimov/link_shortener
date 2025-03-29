package auth

type LoginRequest struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
}
type RegisterResponse struct {
	Token string `json:"token"`
}

type userData struct {
	email    string
	password string
	name     string
}
