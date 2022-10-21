package main

import (
	"os"
	"simple-crud/config"
	"simple-crud/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.DBInit()

	controller := controllers.New(db)

	r := gin.Default()

	r.GET("/person/:id", controller.GetPerson)

	r.Run(":" + os.Getenv("PORT"))
}
