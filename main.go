package main

import (
	"simple-crud/config"
	"simple-crud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := config.DBInit()

	controller := controllers.New(db)

	r := gin.Default()

	r.GET("/person/:id", controller.GetPerson)

	r.Run()
}
