package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Trades is a struct
type Trades struct {
	gorm.Model
	ID            int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Ticket        string
	UID           uint
	Amount        float64
	Profit        float64
	PercentProfit int
	Command       int
	Symbol        string
	CreatedAt     string `sql:"null"`
	UpdatedAt     string `sql:"null"`
}

// GetTrades return user trades by UID
func GetTrades(UID uint) []*Trades {
	trades := make([]*Trades, 0)
	err := GetDB().Table("trades").Where("uid = ?", UID).Find(&trades).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return trades
}
