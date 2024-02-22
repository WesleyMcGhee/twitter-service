package database

import (
	"os"

	"anime-twitter-service/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const Db = *gorm.Database

func DB() *gorm.DB {
  db, err := gorm.Open(postgres.Open(os.Getenv("DB_STRING")), &gorm.Config{})
  if err != nil {
    panic("Failed to connect")
  }

  db.AutoMigrate(&models.User{})

  return db  
}
