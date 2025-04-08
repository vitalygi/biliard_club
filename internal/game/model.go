package game

import (
	"biliard_club/internal/table"
	"biliard_club/internal/user"
	"gorm.io/gorm"
	"time"
)

type Game struct {
	gorm.Model
	TableID   uint
	Table     table.Table `json:"table" gorm:"foreignKey:TableID"`
	UserID    uint
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Status    string    `json:"status"`
	GameType  string    `json:"gameType"`
	TotalCost float64   `json:"totalCost"`
	IsPaid    bool      `json:"isPaid" default:"false"`
}
