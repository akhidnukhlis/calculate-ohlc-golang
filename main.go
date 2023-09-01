package main

import (
	service "calculate-ohlc-golang/services"
	"fmt"
)

func main() {
	s := service.NewService()
	result := s.CalculateOhlcService()

	fmt.Println("Code: ", result.Code)
	fmt.Println("Message: ", result.Message)
	fmt.Println("Data: ", result.Data)
}
