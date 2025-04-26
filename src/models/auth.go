package models

import "github.com/golang-jwt/jwt/v5"

type Credentials struct {
	CPF      string `json:"cpf" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	jwt.RegisteredClaims
}
