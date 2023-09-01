package service

import (
	"calculate-ohlc-golang/common"
	helper "calculate-ohlc-golang/helpers"
	"fmt"
	"net/http"
)

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) CalculateOhlcService() common.Response {
	rawDatas, err := helper.LoadNdjsonFiles()
	if err != nil {
		return common.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Load ndjson files got error",
			Errors:  err,
		}
	}

	var orderDatas []helper.OrderData
	for _, data := range rawDatas {
		orderDatas = append(orderDatas, helper.OrderData{
			StockCode:   data.StockCode,
			OrderNumber: data.OrderNumber,
			Type:        data.Type,
			Quantity:    data.Quantity,
			Price:       data.Price,
		})
	}

	mapOhlc := helper.CalculateOHLC(orderDatas)

	var messages []string
	for stockCode, ohlc := range mapOhlc {
		m := fmt.Sprintf("stock_code: %v, prev: %v, open: %v, low: %v, high: %v, close: %v, volume: %v, value: %v, avg: %v\n",
			stockCode, ohlc.Prev, ohlc.Open, ohlc.Low, ohlc.High, ohlc.Close, ohlc.Volume, ohlc.Value, ohlc.Avg)

		messages = append(messages, m)
	}

	return common.Response{
		Code:    http.StatusOK,
		Message: "Success Calculate OHLC",
		Data:    messages,
	}
}
