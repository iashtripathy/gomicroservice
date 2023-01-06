package serviceModel

import (
	"time"
)

type StockStructure struct {
	StockName    string
	StockSymbol  string
	CurrentPrice float64
	F52WeekHigh  float64
	F52WeekLow   float64
	Timestamp    time.Time
}
