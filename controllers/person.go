package controllers

import (
	"net/http"
	"simple-crud/models"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) GetPerson(ctx *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := ctx.Param("id")

	if err := c.masterDB.Where("id=?", id).First(&person).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	ctx.JSON(http.StatusOK, result)
}
