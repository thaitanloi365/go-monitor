package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/errs"
	"github.com/thaitanloi365/go-monitor/helper"
	"github.com/thaitanloi365/go-monitor/models"
)

// Login Login
// @Tags Admin-Authorization
// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Param data body models.LoginForm true "User Data"
// @Success 200 {object} models.LoginResponse
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/login [post]
func Login(c echo.Context) error {
	var form models.LoginForm
	var cc = c.(*models.CustomContext)

	var err = cc.BindAndValidate(&form)
	if err != nil {
		return err
	}

	user, err := form.Validate()
	if err != nil {
		return err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return err
	}

	var lastLogin = helper.NowIn(user.Timezone).Unix()
	var update = models.User{
		LastLogin:   &lastLogin,
		LoggedOutAt: nil,
		TokenIssuer: user.TokenIssuer,
	}

	err = cc.DB.Model(&user).Update(update).Error
	if err != nil {
		return err
	}

	err = cc.DB.First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errs.ErrUserNotFound
		}
	}
	var response = models.LoginResponse{User: &user, Token: token}

	return cc.Success(response)

}
