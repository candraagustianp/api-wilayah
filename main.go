package main

import (
	"api-wilayah/config"
	"api-wilayah/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.InitConfig()

	//database connect
	database.ConnectDB(config)

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(config.Timeout),
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.Listen(":" + config.Port)
}
