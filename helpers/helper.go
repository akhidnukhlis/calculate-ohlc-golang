package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type RawData struct {
	Type             string `json:"type"`
	OrderNumber      string `json:"order_number"`
	OrderVerb        string `json:"order_verb"`
	Quantity         string `json:"quantity"`
	ExecutedQuantity string `json:"executed_quantity"`
	OrderBook        string `json:"order_book"`
	Price            string `json:"price"`
	ExecutionPrice   string `json:"execution_price"`
	StockCode        string `json:"stock_code"`
}

type OrderData struct {
	StockCode   string `json:"stock_code"`
	OrderNumber string `json:"order_number"`
	Type        string `json:"type"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

type OHLC struct {
	StockCode string
	Prev      int64
	Open      int64
	High      int64
	Low       int64
	Close     int64
	Volume    int64
	Value     int64
	Avg       int64
}

func CalculateOHLC(orderData []OrderData) map[string]OHLC {
	var (
		ohlcMap      = make(map[string]OHLC)
		prevPrice, _ = strconv.ParseInt("0", 10, 64)
		openPrice    = false
	)

	for _, data := range orderData {
		stockCode := data.StockCode
		price, _ := strconv.ParseInt(data.Price, 10, 64)

		if _, exists := ohlcMap[stockCode]; !exists {
			ohlcMap[stockCode] = OHLC{StockCode: stockCode, Prev: price, Open: -1, High: price, Low: price}
		}

		ohlc := ohlcMap[stockCode]
		quantity, _ := strconv.ParseInt(data.Quantity, 10, 64)

		if !openPrice && ohlc.Open == -1 && quantity > 0 && price > prevPrice && data.Type != "A" {
			ohlc.Open = price
		}

		if price > ohlc.High {
			ohlc.High = price
		}

		if price < ohlc.Low || ohlc.Low == 0 {
			ohlc.Low = price
		}

		if quantity > 0 {
			ohlc.Close = price
		}

		if data.Type == "E" || data.Type == "P" {
			ohlc.Volume = quantity
			ohlc.Value += price * quantity
			ohlc.Avg = ohlc.Value / ohlc.Volume
		}

		if data.Type == "A" && quantity <= 1 {
			ohlc.Prev = price
		}

		ohlcMap[stockCode] = ohlc
	}

	return ohlcMap
}

func LoadNdjsonFiles() ([]RawData, error) {
	var (
		directory = "./data/subsetdata"
		rawDatas  []RawData
	)

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("error reading file from directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ndjson" {
			filePath := filepath.Join(directory, file.Name())
			rawDatas, err = processFile(filePath)
			if err != nil {
				return nil, err
			}
		}
	}

	return rawDatas, nil
}

func processFile(filepath string) ([]RawData, error) {
	var rawDatas []RawData

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var rawData RawData
		if err := json.Unmarshal([]byte(line), &rawData); err != nil {
			return nil, fmt.Errorf("error parsing json file: %s", err)
			continue
		}

		q, _ := strconv.Atoi(rawData.ExecutedQuantity)
		if q > 0 {
			rawData.Quantity = rawData.ExecutedQuantity
		}

		p, _ := strconv.Atoi(rawData.ExecutionPrice)
		if p > 0 {
			rawData.Price = rawData.ExecutionPrice
		}

		rawDatas = append(rawDatas, RawData{
			StockCode:   rawData.StockCode,
			OrderNumber: rawData.OrderNumber,
			Type:        rawData.Type,
			Quantity:    rawData.Quantity,
			Price:       rawData.Price,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return rawDatas, nil
}
