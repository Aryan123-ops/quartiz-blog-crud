package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:Aryan@26@/quartiz?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Blog{})
	if err != nil {
		log.Println("Failed to migrate database!")
	}

	DB = database
}
