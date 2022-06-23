package dto

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
)

// Login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	model.User
}

// Register
type AuthRegisterRequest struct {
	model.User
}
type AuthRegisterResponse struct {
	model.User
}
