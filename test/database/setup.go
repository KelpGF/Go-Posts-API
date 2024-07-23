package database

import (
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Post{})

	return db
}
