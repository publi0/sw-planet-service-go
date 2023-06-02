package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("planets.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
