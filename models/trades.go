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

// GetTradesByFilter return user trades by custome filter
func GetTradesByFilter(filter map[string]interface{}) []*Trades {
	trades := make([]*Trades, 0)
	query := GetDB().Table("trades")

	if _, ok := filter["id"]; ok {
		query = query.Where("id = ?", filter["id"])
	}
	if _, ok := filter["uid"]; ok {
		query = query.Where("uid = ?", filter["uid"])
	}

	err := query.Find(&trades).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return trades
}

// GetTradesByUID return user trades by UID
func GetTradesByUID(uid uint) []*Trades {
	trades := make([]*Trades, 0)
	err := GetDB().Table("trades").Where("uid = ?", uid).Find(&trades).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return trades
}

// GetTradeByID return user trade by ID
func GetTradeByID(id int) *Trades {
	trade := &Trades{}
	err := GetDB().Table("trades").First(trade, id).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return trade
}
