package controller

import (
	"api-wilayah/database"
	"api-wilayah/model"
	"api-wilayah/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := &model.User{}

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		var err error
		user.Password, err = util.HashPwd(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		if err := database.SaveData(db, user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error": false,
			"msg":   "user signed up",
		})
	}
}

func Login(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := &model.User{}

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		pass := user.Password

		if err := database.GetAllWhere(db, user, "username = '"+user.Username+"'"); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		if !util.CheckPassHash(pass, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "password incorrect",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"msg":   "login success",
		})
	}
}
