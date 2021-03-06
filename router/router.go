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

	router.Get("/kota/:prov", controller.GetAllKota(db))

	router.Get("/kecamatan/:kota", controller.GetAllKecamatan(db))

	router.Get("/kelurahan/:kec", controller.GetAllKelurahan(db))

	router.Get("/spasial/:spasial", controller.GetSpasialData(db))

	router.Post("/user", controller.CreateUser(db))

	router.Post("/user/login", controller.Login(db))

}
