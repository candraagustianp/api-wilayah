package main

import (
	"api-wilayah/app"
	"api-wilayah/config"
)

func main() {
	config := config.InitConfig()
	app.Run(config)
}
