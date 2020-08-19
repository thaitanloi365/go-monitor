package models

import (
	"net"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/thaitanloi365/go-logging"
)

// JobHealthCheckStatus status
type JobHealthCheckStatus string

// All status
var (
	JobHealthCheckStatusUp   JobHealthCheckStatus = "up"
	JobHealthCheckStatusDown JobHealthCheckStatus = "down"
)

// JobHealthCheck job
type JobHealthCheck struct {
	Tag      string               `gorm:"primary_key" json:"tag"`
	NextAt   time.Time            `json:"next_at"`
	StartAt  string               `json:"start_at"`
	Endpoint string               `json:"endpoint"`
	Timeout  uint64               `json:"timeout"`
	Interval uint64               `json:"interval"`
	Status   JobHealthCheckStatus `gorm:"default:'up'" json:"status"`
	Logs     []*JobHealthCheckLog `gorm:"foreignkey:tag;association_foreignkey:tag" json:"logs"`
}

// HealthcheckJobCreateForm form
type HealthcheckJobCreateForm struct {
	Tag      string `json:"tag" validate:"required" example:"api_healthcheck"`
	Endpoint string `json:"endpoint" validate:"startswith=http" example:"http://localhost:8080"`
	Interval uint64 `json:"interval" validate:"required" example:"30"`
	Timeout  uint64 `json:"timeout" validate:"required,ltfield=Interval" example:"20"`
}

// CreateOrUpdate create or update
func (j JobHealthCheck) CreateOrUpdate() error {
	var job JobHealthCheck
	var err = dbInstance.First(&job, JobHealthCheck{Tag: j.Tag}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = dbInstance.Create(&j).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	err = dbInstance.Model(&job).Update(&j).Error
	return err
}

// HeathCheckJobHandler handler
func HeathCheckJobHandler(tag string, endPoint string, timeout time.Duration) error {
	logging.Global().Infof("Handle job tag = %s endpoint = %s timeout = %d\n", tag, endPoint, timeout)
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: timeout,
		}).Dial,
		TLSHandshakeTimeout: timeout,
	}
	var netClient = &http.Client{
		Timeout:   timeout,
		Transport: netTransport,
	}

	resp, err := netClient.Get(endPoint)
	var log = JobHealthCheckLog{
		Tag:           tag,
		Endpoint:      endPoint,
		StatusCode:    resp.StatusCode,
		StatusMessage: resp.Status,
	}

	log.CreateOrUpdate()

	if log.StatusCode >= 400 {
		var job = JobHealthCheck{
			Tag: tag,
		}
		dbInstance.Model(&job).Update(&JobHealthCheck{Status: JobHealthCheckStatusDown})
	}
	return err
}
