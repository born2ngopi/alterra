package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
