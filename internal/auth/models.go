package auth

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type RegisterResponse struct {
	UserID int64 `json:"userID"`
}

type LoginRequest struct {
	UserID   int64  `json:"userID"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
