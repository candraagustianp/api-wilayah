package controller

import (
	"api-wilayah/database"
	"api-wilayah/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllProvinsi(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//intance to hold json received
		provinsi := []model.Provinsi{}

		if err := database.GetAll(db, &provinsi); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err,
			})
		} else {
			return c.JSON(fiber.Map{
				"error": false,
				"msg":   "success to get all data",
				"data":  provinsi,
			})
		}

	}

}

func CreateProvinsi(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		prov := &model.Kecamatan{}

		if err := c.BodyParser(prov); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		if err := database.SaveData(db, prov); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error": false,
			"msg":   "success save data",
			"data":  prov,
		})
	}
}
