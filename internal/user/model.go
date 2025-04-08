package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name             string    `json:"name"`
	Phone            string    `json:"phone" gorm:"unique;default:''"`
	RegistrationTime time.Time `json:"registrationTime" `
	Discount         float32   `json:"discount" gorm:"default:0"`
	Password         string    `json:"password"`
}
