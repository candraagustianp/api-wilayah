package controller

import (
	"api-wilayah/database"
	"api-wilayah/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllKecamatan(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//intance to hold json received
		kota := []model.Kecamatan{}

		if err := database.GetAllWhere(db, &kota, "id_kab = 3327"); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		} else {
			return c.JSON(fiber.Map{
				"error": false,
				"msg":   "success to get all data",
				"data":  kota,
			})
		}

	}

}
