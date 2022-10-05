package config

import (
	"fmt"
	"simple-crud/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/gorm-crud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db: %+v", err))
	}

	db.AutoMigrate(models.Person{})
	return db
}
