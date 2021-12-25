package router

import (
	"api-wilayah/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetRouting(app *fiber.App, db *gorm.DB) {
	router := app.Group("/api")
	PublicRouting(router, db)
}

func PublicRouting(router fiber.Router, db *gorm.DB) {
	router.Get("/provinsi", controller.GetAllProvinsi(db))

	router.Get("/kota", controller.GetAllKota(db))

	router.Get("/kecamatan", func(c *fiber.Ctx) error {
		return c.SendString("get kecamatan")
	})

	router.Get("/kelurahan", func(c *fiber.Ctx) error {
		return c.SendString("get kelurahan")
	})
}
