package models

// Model is a base model that defines a primary key and timestamp on operation.
type Model struct {
	ID        string `gorm:"primary_key" json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `sql:"index" json:"deleted_at,omitempty"`
}

// HiddenModel is a base model that defines a primary key and timestamp on operation.
type HiddenModel struct {
	ID        string `gorm:"primary_key" json:"id"`
	CreatedAt int64  `json:"-"`
	UpdatedAt int64  `json:"-"`
	DeletedAt *int64 `sql:"index" json:"-"`
}

// NotificationPushCallback response
type NotificationPushCallback struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

// M map
type M map[string]interface{}
