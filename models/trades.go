package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
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

	query = applyFilter(query, filter)

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

func applyFilter(q *gorm.DB, f map[string]interface{}) *gorm.DB {
	if _, ok := f["id"]; ok {
		q = q.Where("id = ?", f["id"])
	}
	if _, ok := f["uid"]; ok {
		q = q.Where("uid = ?", f["uid"])
	}
	if _, ok := f["uids"]; ok {
		q = q.Where("uid IN (?)", f["uids"])
	}
	if _, ok := f["symbol"]; ok {
		q = q.Where("symbol = ?", f["symbol"])
	}
	if _, ok := f["symbols"]; ok {
		q = q.Where("symbols IN (?)", f["symbols"])
	}
	if _, ok := f["dateFrom"]; ok {
		q = q.Where("created_at >= ?", f["dateFrom"])
	}
	if _, ok := f["dateTo"]; ok {
		q = q.Where("created_at <= ?", f["dateTo"])
	}

	return q
}
