package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/thaitanloi365/go-monitor/config"
)

// DB db
type DB struct {
	*gorm.DB
	config *config.Configuration
}

var dbInstance *DB

var models = []interface{}{
	&User{},
}

// Setup bootstrap app
func Setup() {
	SetupDB()
	SetupRelation()

	seeds()
}

// SetupRelation db relation
func SetupRelation() {
	dbInstance.AutoMigrate(
		models...,
	)

}

// SetupDB Setup database
func SetupDB() *DB {
	fmt.Println("[================ Setup database ================]")
	var cfg = config.GetInstance()

	db, err := gorm.Open("sqlite3", cfg.DBPath)
	if err != nil {
		panic("failed to connect database")
	}

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	gorm.DefaultCallback.Create().Before("gorm:save_before_associations").Register("app:update_xid_when_create", updateIDForCreateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	dbInstance = &DB{
		db.Debug(),
		cfg,
	}

	return dbInstance
}

// GetDBInstance get db instance
func GetDBInstance() *DB {
	if dbInstance == nil {
		panic("Must be call SetupDB first")
	}
	return dbInstance
}

func seeds() {
	seedAccounts()
}
