package server

import (
	"go-clean-todo/infrastructure/mysql"
	"go-clean-todo/presentation/settings"
	"go-clean-todo/server/route"
	"log"

	"github.com/joho/godotenv"
)

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run() {
	loadENV()
	mysql.Connect()
	mysql.RunMigration()

	api := settings.NewGinEngine()
	route.InitRoute(api)

	api.Run()
}
