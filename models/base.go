// Package models ...
package models

import (
	"go-echo-vue/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	validator "gopkg.in/go-playground/validator.v9"
)

var db = DB()
var validate = validator.New()

// DB connect database. setting logging.
func DB() *gorm.DB {
	db, err := gorm.Open(config.DbType, config.DbURL)
	if err != nil {
		panic("failed to connect database")
	}

	logPath := config.ProjectPath + "logs/gorm.log"
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	db.SingularTable(true)
	db.LogMode(true)
	db = db.Debug()
	db.SetLogger(log.New(file, "", 0))

	return db
}
