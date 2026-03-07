package auth

type LoginResopnse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name string `json:"name" validate:"required,username_valid"`
}

type RegisterResponse struct {
	Email    string `json:"email"`
	Name string `json:"name"`
}