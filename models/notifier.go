package models

// NotifierProvider provider
type NotifierProvider string

// All providers supported
var (
	NotifierProviderSlack    NotifierProvider = "slack"
	NotifierProviderSMTP     NotifierProvider = "smtp"
	NotifierProviderSendgrid NotifierProvider = "sendgrid"
)

// Notifier notifier
type Notifier struct {
	Provider NotifierProvider `gorm:"primary_key" json:"provider"`
	Disabled *bool            `gorm:"default:'false'" json:"disabled"`
	Metadata JSONRaw          `json:"-"`
}
