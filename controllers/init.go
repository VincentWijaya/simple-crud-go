package controllers

import (
	"github.com/jinzhu/gorm"
)

type Controllers struct {
	masterDB *gorm.DB
}

func New(db *gorm.DB) *Controllers {
	return &Controllers{
		masterDB: db,
	}
}
