package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model

	Name     string `json:"name" gorm:"size:200;not null"`
	Email    string `json:"email" gorm:"size:200;not null;unique"`
	Password string `json:"password,omitempty"`
}

// BeforeCreate is a method for struct User
// gorm call this method before they execute query
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.HashPassword()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

// HashPassword is a method for struct User for Hashing password
func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}

// GenerateToken is a method for struct User for creating new jwt token
func (u *User) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // we set expired in 72 hour
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
