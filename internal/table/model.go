package table

import (
	"biliard_club/internal/user"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Type              string  `json:"type"`
	PriceBeforeSwitch float32 `json:"priceBeforeSwitch"`
	PriceAfterSwitch  float32 `json:"priceAfterSwitch"`
	SwitchTime        int     `json:"switchTime"`
	SwitchLong        int     `json:"switchLong"`
	UserID            uint
	User              user.User `json:"user" gorm:"foreignKey:UserID"`
}
