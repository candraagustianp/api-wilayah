package app

import (
	"api-wilayah/config"
	"api-wilayah/database"
	"api-wilayah/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Run(config *config.Config) {
	db := connectDatabase(config)
	// migrateTabel(db)
	runServer(config, db)
}

func runServer(conf *config.Config, db *gorm.DB) {
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(conf.Timeout),
	})
	router.GetRouting(app, db)

	app.Listen(":" + config.GetString("PORT"))
}

func connectDatabase(config *config.Config) *gorm.DB {
	db := database.ConnectDB(config)
	return db
}

// func migrateTabel(db *gorm.DB) {
// 	database.AutoMigrate(db, "provinsi", &model.Provinsi{})
// 	database.AutoMigrate(db, "kab_kotas", &model.KabKotas{})
// 	database.AutoMigrate(db, "kecamatan", &model.Kecamatan{})
// 	database.AutoMigrate(db, "kelurahan", &model.Kelurahan{})
// }
