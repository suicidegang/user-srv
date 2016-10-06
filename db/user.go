package db

import (
	"github.com/jinzhu/gorm"

	"time"
)

type User struct {
	gorm.Model
	Email            string    `gorm:"unique_index"`
	EmailConfirmedAt time.Time `gorm:"null"`
	Password         string
	Name             string
	Username         string
}
