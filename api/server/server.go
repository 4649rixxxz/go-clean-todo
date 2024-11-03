package server

import (
	"go-clean-todo/infrastructure/mysql"
	"go-clean-todo/server/route"
	"log"

	"github.com/gin-gonic/gin"
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

	api := gin.Default()
	route.InitRoute(api)

	api.Run()
}
