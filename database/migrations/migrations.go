package migrations

import (
	"github.com/reafreitas1/booksAPI_Golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
