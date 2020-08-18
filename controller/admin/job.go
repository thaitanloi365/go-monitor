package controllers

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/models"
	"github.com/thaitanloi365/go-monitor/scheduler"
)

// GetListJobHealthCheck Get list scheduled health check jobs
// @Tags Admin-Job-HealthCheck
// @Summary Get list scheduled health check jobs
// @Description Get list scheduled health check jobs
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job_healthcheck/list [get]
func GetListJobHealthCheck(c echo.Context) error {
	var cc = c.(*models.CustomContext)

	var jobs []*models.JobHealthCheck
	var err = cc.DB.Preload("Logs").Find(&jobs).Error
	if err != nil {
		return err
	}

	for _, job := range scheduler.GetInstance().Jobs() {
		var j = models.JobHealthCheck{
			Tag:     job.Tags()[0],
			StartAt: job.ScheduledAtTime(),
			NextAt:  job.ScheduledTime(),
		}

		for _, storedJob := range jobs {
			if j.Tag == storedJob.Tag {
				cc.DB.Model(storedJob).Update(&j)
			} else {
				cc.DB.Model(storedJob).Delete(&models.JobHealthCheck{})
			}
		}
	}
	return cc.Success(jobs)

}

// RemoveJobHealthCheckByTag Remove job by tag
// @Tags Admin-Job
// @Summary Remove job by tag
// @Description Remove job by tag
// @Param tag path string true "Tag of job"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job_healthcheck/{tag} [delete]
func RemoveJobHealthCheckByTag(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var tag = cc.GetPathParamString("tag")

	scheduler.GetInstance().RemoveJobByTag(tag)
	cc.DB.Delete(&models.JobHealthCheck{}, "tag = ?", tag)

	return cc.Success(fmt.Sprintf("Job tag = %s was deleted", tag))

}

// GetJobHealthCheckByTag Remove job by tag
// @Tags Admin-Job
// @Summary Remove job by tag
// @Description Remove job by tag
// @Param tag path string true "Tag of job"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job_healthcheck/{tag} [delete]
func GetJobHealthCheckByTag(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var tag = cc.GetPathParamString("tag")
	var job models.JobHealthCheck
	for _, j := range scheduler.GetInstance().Jobs() {
		if j.Tags()[0] == tag {
			cc.DB.First(&job, "tag = ?", tag)
		} else {
			cc.DB.Delete(&models.JobHealthCheck{}, "tag = ?", tag)
		}
	}

	return cc.Success(fmt.Sprintf("Job tag = %s was deleted", tag))

}

// AddJobHealthCheck Add healthcheck job
// @Tags Admin-Job
// @Summary Add healthcheck job
// @Description Add healthcheck job
// @Param data body models.HealthcheckJobCreateForm true "Form"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job_healthcheck [post]
func AddJobHealthCheck(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var form models.HealthcheckJobCreateForm
	var err = cc.BindAndValidate(&form)
	if err != nil {
		return err
	}

	err = scheduler.GetInstance().RemoveJobByTag(form.Tag)
	if err != nil {
		fmt.Println(err)
	}

	job, err := scheduler.GetInstance().
		Every(form.Interval).
		Seconds().
		SetTag([]string{form.Tag}).
		Do(models.HeathCheckJobHandler, form.Endpoint, time.Duration(form.Timeout)*time.Second)
	if err != nil {
		return err
	}

	var j = models.JobHealthCheck{
		Tag:      job.Tags()[0],
		StartAt:  job.ScheduledAtTime(),
		NextAt:   job.ScheduledTime(),
		Endpoint: form.Endpoint,
		Interval: form.Interval,
		Timeout:  form.Timeout,
	}

	err = j.CreateOrUpdate()
	if err != nil {
		return err
	}

	return cc.Success(j)

}
