package services

import (
	"github.com/d3z41k/jsonrpc-server-boilerplate/models"
)

type Filter struct {
	ID       uint
	UID      uint
	UIDS     []uint
	Symbol   string
	Symbols  []string
	DateFrom string
	DateTo   string
}

type TradesService struct{}

func (ft *Filter) getFilter() map[string]interface{} {
	f := make(map[string]interface{})

	if ft.ID != 0 {
		f["id"] = ft.ID
	}
	if ft.UID != 0 {
		f["uid"] = ft.UID
	}
	if len(ft.UIDS) > 0 {
		f["uids"] = ft.UIDS
	}
	if ft.Symbol != "" {
		f["symbol"] = ft.Symbol
	}
	if len(ft.Symbols) > 0 {
		f["symbols"] = ft.Symbols
	}
	if ft.DateFrom != "" {
		f["dateFrom"] = ft.DateFrom
	}
	if ft.DateTo != "" {
		f["dateTo"] = ft.DateTo
	}

	return f
}

func (ts *TradesService) GetCountTrades(in *Filter, out *int) error {
	// fmt.Printf("call getCountTrades %v", in)

	data := models.GetTradesByFilter(in.getFilter())

	*out = len(data)
	return nil
}
