package main

import (
	"github.com/KelpGF/Go-Posts-API/configs"
	_ "github.com/KelpGF/Go-Posts-API/docs"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure"
	"github.com/KelpGF/Go-Posts-API/internal/presentation"
)

// @title           Go Posts API
// @version         1.0
// @description     Posts API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Kelvin Gomes
// @contact.url    https://www.linkedin.com/in/kelvin-gomes-fernandes
// @contact.email  kelvingomesdeveloper@gmail.com

// @host      localhost:3000
// @BasePath  /
func main() {
	config := configs.NewConfig()

	db := infrastructure.StartDatabase(config)
	defer infrastructure.CloseDatabase()

	presentation.StartWebServer(db, config)
}
