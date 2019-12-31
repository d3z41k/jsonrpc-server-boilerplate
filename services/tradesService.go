package services

import (
	"fmt"

	"github.com/d3z41k/jsonrpc-server-boilerplate/models"
)

type Filter struct {
	UID uint
}

type TradesService struct{}

func (ts *TradesService) GetCountTrades(in *Filter, out *int) error {
	fmt.Println("call getCountTrades", in)

	filter := make(map[string]interface{})

	filter["uid"] = in.UID

	data := models.GetTradesByFilter(filter)

	fmt.Println(len(data))

	*out = len(data)
	return nil
}
