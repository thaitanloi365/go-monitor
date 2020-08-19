package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/models"
	"github.com/thaitanloi365/go-monitor/scheduler"
)

// GetListNotifier Get list notifier
// @Tags Admin-Notifier
// @Summary Get list notifier
// @Description Get list notifier
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/notifier/list [get]
func GetListNotifier(c echo.Context) error {
	var cc = c.(*models.CustomContext)

	var notifiers []*models.Notifier
	var err = cc.DB.Find(&notifiers).Error
	if err != nil {
		return err
	}
	return cc.Success(notifiers)

}

// UpdateNotifier Update notifier
// @Tags Admin-Notifier
// @Summary Update notifier
// @Description Update notifier
// @Param tag path string true "Tag of job"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/notifier/{provider} [put]
func UpdateNotifier(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var tag = cc.GetPathParamString("tag")

	scheduler.GetInstance().RemoveJobByTag(tag)
	cc.DB.Delete(&models.JobHealthCheck{}, "tag = ?", tag)

	return cc.Success(fmt.Sprintf("Job tag = %s was deleted", tag))

}

// GetNotifier Get notifier
// @Tags Admin-Notifier
// @Summary Get notifier
// @Description Get notifier
// @Param tag path string true "Tag of job"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/notifier/{provider} [get]
func GetNotifier(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var tag = cc.GetPathParamString("tag")

	scheduler.GetInstance().RemoveJobByTag(tag)
	cc.DB.Delete(&models.JobHealthCheck{}, "tag = ?", tag)

	return cc.Success(fmt.Sprintf("Job tag = %s was deleted", tag))

}
