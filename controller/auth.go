package controller

import (
	"time"

	"github.com/brianvoe/sjwt"
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/models"
)

// Login Login
// @Tags Authorization
// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Param data body models.LoginForm true "User Data"
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/login [post]
func Login(c echo.Context) error {
	var form models.LoginForm
	var cc = c.(*models.CustomContext)

	var err = cc.BindAndValidate(&form)
	if err != nil {
		cc.Logging.Error(err)
		return err
	}

	var claims = sjwt.New()
	claims.SetExpiresAt(time.Now().Add(cc.Config.JWTExpiry))

	var token = claims.Generate([]byte(cc.Config.JWTSecret))

	return cc.Success(token)

}
