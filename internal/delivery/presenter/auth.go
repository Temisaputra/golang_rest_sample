package presenter

type RegisterRequest struct {
	Username string `json:"username" validate:"required,not_blank"`
	Role     string `json:"role" validate:"required"`
	Email    string `json:"email" validate:"required,email,not_blank"`
	Password string `json:"password" validate:"required,min=6,not_blank"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,not_blank"`
	Password string `json:"password" validate:"required,min=6,not_blank"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}
