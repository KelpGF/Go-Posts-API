package main

import (
	"github.com/KelpGF/Go-Posts-API/configs"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure"
	"github.com/KelpGF/Go-Posts-API/internal/presentation"
)

func main() {
	config := configs.NewConfig()

	db := infrastructure.StartDatabase(config)
	defer infrastructure.CloseDatabase()

	presentation.StartWebServer(db, config)
}
