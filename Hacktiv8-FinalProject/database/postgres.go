package database

import (
	"finalProject/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST   = "172.17.0.3"
	POSTGRES_PORT   = "5432"
	POSTGRES_USER   = "postgres"
	POSTGRES_PASS   = "postgres"
	POSTGRES_DBNAME = "Hacktiv8-FinalProjectGO"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASS,
		POSTGRES_DBNAME,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	return db
}

// func GetDB() *gorm.DB {
// 	return db
// }
