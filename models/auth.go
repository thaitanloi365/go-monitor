package models

import "github.com/dgrijalva/jwt-go"

// LoginForm login form
type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// JwtClaims custom claims
type JwtClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
