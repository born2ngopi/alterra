package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Model

	Name        string `json:"name" gorm:"size:200;unique;not null"`
	Stock       int    `json:"stock" gorm:"not null"`
	Description string `json:"description"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}
