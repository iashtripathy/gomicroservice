package client

import (
	"ModularMicroservice/service/serviceModel"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/piquette/finance-go/quote"
)

func GetStocks(s string, channel chan serviceModel.StockStructure) {
	fmt.Println("Making Api call here")
	smbl := s
	q, err := quote.Get(smbl)
	if err != nil {
		log.Fatal(err)
	}
	//var index int = len(s)
	//index = index + 1
	var temp serviceModel.StockStructure
	temp.StockName = q.ShortName
	temp.StockSymbol = smbl
	temp.CurrentPrice = q.Ask
	temp.F52WeekHigh = q.FiftyTwoWeekHigh
	temp.F52WeekLow = q.FiftyTwoWeekLow
	temp.Timestamp = time.Now()

	channel <- temp
	//abc = append(abc, temp)
	fmt.Println(reflect.TypeOf(q))
	//return abc
}
