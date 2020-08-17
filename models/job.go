package models

import "time"

// Job job
type Job struct {
	Tags    []string  `json:"tags"`
	NextAt  time.Time `json:"next_at"`
	StartAt string    `json:"start_at"`
}

// HealthcheckJobCreateForm form
type HealthcheckJobCreateForm struct {
	Tag      string `json:"tag" validate:"required" example:"api_healthcheck"`
	Endpoint string `json:"endpoint" validate:"startswith=http" example:"http://localhost:8080"`
	Interval uint64 `json:"interval" validate:"required" example:"30"`
	Timeout  uint64 `json:"timeout" validate:"required,ltfield=Interval" example:"20"`
}
