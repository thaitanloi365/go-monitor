package models

// JobHealthCheckLog log
type JobHealthCheckLog struct {
	ID            string `gorm:"primary_key" json:"id"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	Tag           string `json:"tag"`
	Endpoint      string `json:"endpoint"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	CheckedAt     int64  `json:"checked_at"`
}
