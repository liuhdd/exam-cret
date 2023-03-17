package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sync"
)

var dbInstance *gorm.DB
var dbOnce sync.Once

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		dbInstance = db
	})
	return dbInstance
}
