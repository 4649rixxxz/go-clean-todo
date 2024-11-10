package mysql

import (
	"time"
)

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Todos     []Todo `gorm:"foreignKey:UserID"`
}

type Todo struct {
	TodoID           uint       `gorm:"primaryKey"`
	UserID           uint       `gorm:"not null"`
	Title            string     `gorm:"not null"`
	Description      string     `gorm:"not null"`
	AttachedFilePath *string    `gorm:"default: null"`
	CompletedAt      *time.Time `gorm:"default: null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
