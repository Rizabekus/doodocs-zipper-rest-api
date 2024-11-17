package app

import (
	"log"

	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/handlers"
	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/services"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/loggers"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/utils"
	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	loggers.InitLoggers()
	file, line, _ := utils.GetCallerInfo()
	loggers.InfoLog(file, line, "Loaded the configuration data from .env")

	service := services.ServiceInstance()
	handler := handlers.HandlersInstance(service)

	Routes(handler)
}
