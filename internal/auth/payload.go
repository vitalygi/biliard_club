package auth

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
type RegisterResponse struct {
	Token string `json:"token"`
}
