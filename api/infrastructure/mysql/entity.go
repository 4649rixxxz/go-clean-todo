package mysql

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Todos     []Todo `gorm:"foreignKey:UserID"`
}

type Todo struct {
	TodoID           uint       `gorm:"primaryKey"`
	UserID           uint       `gorm:"not null"`
	title            string     `gorm:"not null"`
	description      string     `gorm:"not null"`
	attachedFilePath *string    `gorm:"default: null"`
	CompletedAt      *time.Time `gorm:"default: null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
