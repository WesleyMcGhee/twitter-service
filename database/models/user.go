package models

import "time"

type User struct {
  ID uint `gorm:"primaryKey"`
  Email string
  Username string
  PasswordHash string
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time `gorm:"index"`
}
