package main

import (
	"GoRestApi/internal/api"
	"GoRestApi/pkg/config"
	"GoRestApi/pkg/data"
)

func main() {
	//
	cfg := config.New()
	db := data.NewMongoConnection(cfg)
	defer db.Disconnect()

	application := api.New()
	application.Start()
}
