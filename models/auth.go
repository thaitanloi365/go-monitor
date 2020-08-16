package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/thaitanloi365/go-monitor/errs"
)

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

// Validate validate user's info for create
func (form LoginForm) Validate() (User, error) {
	var user User

	var err = dbInstance.Find(&user, &User{Email: form.Email}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return user, errs.ErrUserNotFound
		}
		return user, err
	}

	err = user.ComparePassword(form.Password)
	if err != nil {
		return user, errs.ErrPasswordIncorrect
	}

	return user, nil
}
