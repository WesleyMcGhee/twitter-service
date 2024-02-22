package models

import (
	"time"

	"gorm.io/gorm"
)

type Posts struct {
	ID string `gorm:"primaryKey"`
	User string
	Post string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}