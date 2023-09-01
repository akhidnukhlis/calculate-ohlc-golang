package helper_test

import (
	helper "calculate-ohlc-golang/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadNdjsonFiles(t *testing.T) {
	rawDatas, err := helper.LoadNdjsonFiles()

	assert.NotNil(t, rawDatas, "raw data should not be nil")
	assert.NoError(t, err, "raw data should not be error")
}

func TestCalculateOHLC(t *testing.T) {
	orderData := []helper.OrderData{
		{StockCode: "BBRI", Type: "A", Price: "6000", Quantity: "0"},
		{StockCode: "BBRI", Type: "P", Price: "6050", Quantity: "2"},
		{StockCode: "BBRI", Type: "A", Price: "5950", Quantity: "5"},
		{StockCode: "BBRI", Type: "A", Price: "7150", Quantity: "8"},
		{StockCode: "BBRI", Type: "E", Price: "7100", Quantity: "3"},
		{StockCode: "BBRI", Type: "A", Price: "7200", Quantity: "12"},

		{StockCode: "BBCA", Type: "A", Price: "8000", Quantity: "0"},
		{StockCode: "BBCA", Type: "P", Price: "8050", Quantity: "100"},
		{StockCode: "BBCA", Type: "P", Price: "7950", Quantity: "500"},
		{StockCode: "BBCA", Type: "A", Price: "8150", Quantity: "200"},
		{StockCode: "BBCA", Type: "E", Price: "8100", Quantity: "300"},
		{StockCode: "BBCA", Type: "A", Price: "8200", Quantity: "100"},
	}

	ohlcMap := helper.CalculateOHLC(orderData)

	//Test for the first stock code
	if ohlc, exists := ohlcMap["BBRI"]; exists {
		if ohlc.Prev != 6000 {
			t.Errorf("for BBRI, expected Prev to be 6000, but got %v", ohlc.Prev)
		}
		if ohlc.Open != 6050 {
			t.Errorf("for BBRI, expected Open to be 6050, but got %v", ohlc.Open)
		}
		if ohlc.High != 7200 {
			t.Errorf("for BBRI, expected High to be 7200, but got %v", ohlc.High)
		}
		if ohlc.Low != 5950 {
			t.Errorf("for BBRI, expected Low to be 5950, but got %v", ohlc.Low)
		}
		if ohlc.Close != 7200 {
			t.Errorf("for BBRI, expected Close to be 7200, but got %v", ohlc.Close)
		}
	} else {
		t.Errorf("expected BBRI stock code in OHLC map, but not found")
	}

	//Test for the first stock code
	if ohlc, exists := ohlcMap["BBCA"]; exists {
		if ohlc.Prev != 8000 {
			t.Errorf("for BBCA, expected Prev to be 8000, but got %v", ohlc.Prev)
		}
		if ohlc.Open != 8050 {
			t.Errorf("for BBCA, expected Open to be 8050, but got %v", ohlc.Open)
		}
		if ohlc.High != 8200 {
			t.Errorf("for BBCA, expected High to be 8200, but got %v", ohlc.High)
		}
		if ohlc.Low != 7950 {
			t.Errorf("for BBCA, expected Low to be 7950, but got %v", ohlc.Low)
		}
		if ohlc.Close != 8200 {
			t.Errorf("for BBCA, expected Close to be 8200, but got %v", ohlc.Close)
		}
	} else {
		t.Errorf("expected BBCA stock code in OHLC map, but not found")
	}
}
