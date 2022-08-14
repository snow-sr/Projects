package migrations

import (
	"github.com/snow-sr/learningGo/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.TodoItem{})
}
