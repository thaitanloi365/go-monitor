package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/helper"
	"github.com/thaitanloi365/go-monitor/models"
)

// Logout logout
// @Tags Admin-Me
// @Summary Logout
// @Description Logout
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Failure 404 {object} errs.Error
// @Router /api/v1/admin/me/logout [delete]
// @Security ApiKeyAuth
func Logout(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var user models.User

	var err = cc.GetUserFromContext(&user)
	if err != nil {
		return err
	}

	var updates = models.M{
		"token_issuer":  "",
		"logged_out_at": helper.NowIn(user.Timezone).Unix(),
	}

	err = cc.DB.Model(models.User{}).Where("id = ?", user.ID).Update(updates).Error

	return cc.Success("Logout successfully")

}
