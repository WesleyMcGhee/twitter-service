package database

import (
  "os"

  "anime-twitter-service/database/models"


  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

func DB() *gorm.DB {
  db, err := gorm.Open(postgres.Open(os.Getenv("DB_STRIGN")), &gorm.Config{})
  if err != nil {
    panic("Failed to connect")
  }

  db.AutoMigrate(&models.User{})

  return db  
}
