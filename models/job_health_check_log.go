package models

import "github.com/jinzhu/gorm"

// JobHealthCheckLog log
type JobHealthCheckLog struct {
	ID            string `gorm:"primary_key" json:"id"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	Tag           string `json:"tag"`
	Endpoint      string `json:"endpoint"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// CreateOrUpdate create or update
func (j JobHealthCheckLog) CreateOrUpdate() error {
	var jobLog JobHealthCheckLog
	var err = dbInstance.First(&jobLog, JobHealthCheckLog{Tag: j.Tag, StatusCode: j.StatusCode, Endpoint: j.Endpoint}).Error
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

	err = dbInstance.Model(&jobLog).Update(&j).Error
	return err
}
