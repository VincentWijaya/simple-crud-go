package config

import (
	"fmt"
	"os"
	"simple-crud/models"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("POSTGRE_HOST")
	username = os.Getenv("POSTGRE_USERNAME")
	password = os.Getenv("POSTGRE_PASSWORD")
	database = os.Getenv("POSTGRE_DATABASE")
	port     = os.Getenv("POSTGRE_PORT")
)

func DBInit() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db: %+v", err))
	}

	db.AutoMigrate(models.Person{})
	return db
}
