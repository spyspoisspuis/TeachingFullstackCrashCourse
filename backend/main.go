package main

import (
	"backend/config"
	"backend/database"
	"backend/server"
)

func main() {
	config.InitializeViper("./")
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)

	svr := server.NewGinServer(&cfg, &db)
	svr.Start()
}
