package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/models"
	"github.com/thaitanloi365/go-monitor/scheduler"
)

// GetListJob Get list scheduled jobs
// @Tags Admin-Job
// @Summary Get list scheduled jobs
// @Description Get list scheduled jobs
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job/list [get]
func GetListJob(c echo.Context) error {
	var cc = c.(*models.CustomContext)

	var jobs []*models.Job
	for _, job := range scheduler.GetInstance().Jobs() {
		var j = models.Job{
			Tags:    job.Tags(),
			StartAt: job.ScheduledAtTime(),
			NextAt:  job.ScheduledTime(),
		}
		jobs = append(jobs, &j)
	}
	return cc.Success(jobs)

}

// RemoveJobByTag Remove job by tag
// @Tags Admin-Job
// @Summary Remove job by tag
// @Description Remove job by tag
// @Param tag path string true "Tag of job"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job/{tag} [delete]
func RemoveJobByTag(c echo.Context) error {
	var cc = c.(*models.CustomContext)

	var jobs []*models.Job
	for _, job := range scheduler.GetInstance().Jobs() {
		var j = models.Job{
			Tags:    job.Tags(),
			StartAt: job.ScheduledAtTime(),
			NextAt:  job.ScheduledTime(),
		}
		jobs = append(jobs, &j)
	}
	return cc.Success(jobs)

}

// AddHealthcheckJob Add healthcheck job
// @Tags Admin-Job
// @Summary Add healthcheck job
// @Description Add healthcheck job
// @Param data body models.HealthcheckJobCreateForm true "Form"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/job/add_healthcheck [post]
func AddHealthcheckJob(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var form models.HealthcheckJobCreateForm
	var err = cc.BindAndValidate(&form)
	if err != nil {
		return err
	}

	var handlerFunc = func(endPoint string, timeout time.Duration) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		go func() {
			select {
			case <-time.After(timeout * 2):
				fmt.Println("overslept after", timeout*2)
			case <-ctx.Done():
				fmt.Println(ctx.Err())
			}
		}()

		req, err := http.NewRequestWithContext(ctx, "GET", endPoint, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(data))
		return nil
	}

	err = scheduler.GetInstance().RemoveJobByTag(form.Tag)
	if err != nil {
		fmt.Println(err)
	}

	job, err := scheduler.GetInstance().
		Every(form.Interval).
		Seconds().
		SetTag([]string{form.Tag}).
		Do(handlerFunc, form.Endpoint, time.Duration(form.Timeout)*time.Second)
	if err != nil {
		return err
	}

	var j = models.Job{
		Tags:    job.Tags(),
		StartAt: job.ScheduledAtTime(),
		NextAt:  job.ScheduledTime(),
	}
	return cc.Success(j)

}
