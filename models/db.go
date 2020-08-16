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
