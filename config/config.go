package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

type Config struct {
	Port       string
	Timeout    int
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

func InitConfig() *Config {
	if err := gotenv.Load(".env"); err != nil {
		log.Panic(err)
	}
	config := &Config{
		Port:       GetString("PORT"),
		Timeout:    GetInt("TIMEOUT"),
		DBUsername: GetString("DB_USERNAME"),
		DBPassword: GetString("DB_PASSWORD"),
		DBHost:     GetString("DB_HOST"),
		DBPort:     GetInt("DB_PORT"),
		DBName:     GetString("DB_NAME"),
	}

	return config
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	v, err := strconv.Atoi(GetString(key))
	if err != nil {
		log.Fatalf("could not parse key: %s, error: %s", key, err)
	}

	return v
}
