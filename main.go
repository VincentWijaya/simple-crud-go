package main

import (
	"simple-crud/config"
	"simple-crud/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()

	controller := controllers.New(db)

	r := gin.Default()

	r.GET("/person/:id", controller.GetPerson)
}
