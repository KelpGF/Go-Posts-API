package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Post struct {
	gorm.Model

	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"text;not null"`
	Body        string `gorm:"text;not null"`
	AuthorName  string `gorm:"text;not null"`
	PublishedAt string `gorm:"datetime;not null"`
}

func main() {
	log.Println("Starting server...")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	err = db.AutoMigrate(&Post{})
	if err != nil {
		log.Fatal("Failed to run migrations. \n", err)
		os.Exit(2)
	}

	log.Println("migrations ran successfully")

	db.Create(&Post{
		Title:       "Hello World",
		Body:        "This is a test post",
		AuthorName:  "John Doe",
		PublishedAt: "2021-09-01 00:00:00",
	})

	p := Post{}
	db.First(&p, 2)

	log.Printf("ID: %d\n", p.ID)
	log.Printf("Title: %s\n", p.Title)
	log.Printf("Body: %s\n", p.Body)
	log.Printf("Author: %s\n", p.AuthorName)
	log.Printf("Published At: %s\n", p.PublishedAt)
	log.Printf("Created At: %s\n", p.CreatedAt)
	log.Printf("Updated At: %s\n", p.UpdatedAt)
}
