package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/KelpGF/Go-Posts-API/configs"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConnection *gorm.DB

func StartDatabase(config *configs.ConfigType) *gorm.DB {
	dsn := makeDSN(config)

	db := openConnection(dsn)

	runMigrations(db)

	dbConnection = db

	return dbConnection
}

func CloseDatabase() {
	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Fatal("Failed to close database. \n", err)
		os.Exit(2)
	}

	sqlDB.Close()

	log.Println("Database connection closed")
}

func makeDSN(config *configs.ConfigType) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
}

func openConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Database connected")

	return db
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entities.Post{})
	if err != nil {
		log.Fatal("Failed to run migrations. \n", err)
		os.Exit(2)
	}

	log.Println("Migrations run successfully")
}
