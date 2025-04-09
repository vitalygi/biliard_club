package domain

import (
	"gorm.io/gorm"
	"time"
)

type Game struct {
	gorm.Model
	TableID   uint
	Table     Table `json:"table" gorm:"foreignKey:TableID"`
	UserID    uint
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Status    string    `json:"status"`
	GameType  string    `json:"gameType"`
	TotalCost float64   `json:"totalCost"`
	IsPaid    bool      `json:"isPaid" default:"false"`
}
