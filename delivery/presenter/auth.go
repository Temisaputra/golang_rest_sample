package presenter

type RegisterRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        Users  `json:"user"`
}
