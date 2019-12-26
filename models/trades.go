package models

import (
	"fmt"
)

// Trades is a struct
type Trades struct {
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

// GetTradesByUID return user trades by UID
func GetTradesByUID(UID uint) []*Trades {
	trades := make([]*Trades, 0)
	err := GetDB().Table("trades").Where("uid = ?", UID).Find(&trades).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return trades
}

// GetTradeByID return user trade by ID
func GetTradeByID(ID int) *Trades {
	trade := &Trades{}
	err := GetDB().Table("trades").First(trade, ID).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return trade
}
