package main

import (
	"github.com/KelpGF/Go-Posts-API/configs"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure"
	"github.com/KelpGF/Go-Posts-API/internal/presentation"
)

func main() {
	config := configs.NewConfig()

	db := infrastructure.StartDatabase(config)

	presentation.StartWebServer(db, config)
}
