package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

type config struct {
	Port       string
	Timeout    int
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

func InitConfig() *config {
	if err := gotenv.Load(".env"); err != nil {
		log.Panic(err)
	}
	config := &config{
		Port:    GetString("PORT"),
		Timeout: GetInt("TIMEOUT"),
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
