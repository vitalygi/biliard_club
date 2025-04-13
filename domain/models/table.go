package models

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Type              string  `json:"type"`
	PriceBeforeSwitch float32 `json:"priceBeforeSwitch"`
	PriceAfterSwitch  float32 `json:"priceAfterSwitch"`
	SwitchTime        int     `json:"switchTime"`
	SwitchLong        int     `json:"switchLong"`
}
