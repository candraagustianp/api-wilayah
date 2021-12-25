package database

import (
	"api-wilayah/config"
	"fmt"

	log "github.com/siruspen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)))
	if err != nil {
		log.Fatalf("unable to initiate database connection: %v", err)
	}
	log.Infoln("Success initiate database connection")
	return db
}
