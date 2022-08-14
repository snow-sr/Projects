package database

import (
	"log"

	"github.com/snow-sr/learningGo/database/migrations"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	database, err := gorm.Open(sqlite.Open("communism.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Err: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(10)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
