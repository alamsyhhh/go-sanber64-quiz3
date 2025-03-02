package dto

type RegisterRequest struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct {
	Password string `json:"password" example:"test12345" binding:"required"`
	Username string `json:"username" example:"test 3" binding:"required"`
}

type UpdateUserRequest struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}