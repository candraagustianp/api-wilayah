package main

import (
	"api-wilayah/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.InitConfig()

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(config.Timeout),
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.Listen(":" + config.Port)
}
