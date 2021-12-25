package main

import (
	"api-wilayah/config"
	"api-wilayah/database"
	"api-wilayah/router"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.InitConfig()

	//database connect
	db := database.ConnectDB(config)

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(config.Timeout),
	})

	router.GetRouting(app, db)

	app.Listen(":" + config.Port)
}
