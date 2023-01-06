package service

import (
	"ModularMicroservice/client"
	"ModularMicroservice/service/serviceModel"
	"errors"
)

// StringService provides operations on strings.
type StockService interface {
	FetchStocks([]string) ([]serviceModel.StockStructure, error)
	//Count(string) int
}

// function to make api calls

// stringService is a concrete implementation of StringService
type StockServiceStruct struct{}

func (StockServiceStruct) FetchStocks(s []string) ([]serviceModel.StockStructure, error) {

	if len(s) == 0 {
		var emptyStock []serviceModel.StockStructure
		return emptyStock, ErrEmpty
	}
	//Make api call here

	var abc []serviceModel.StockStructure

	st := make(chan serviceModel.StockStructure)
	for _, val := range s {
		//fmt.Println(val)
		go client.GetStocks(val, st)
	}

	for i := 0; i < len(s); i++ {
		abc = append(abc, <-st)
	}
	/*

		url := "https://yh-finance.p.rapidapi.com/auto-complete?q=tesla&region=US"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("X-RapidAPI-Key", "****")
		req.Header.Add("X-RapidAPI-Host", "yh-finance.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))
	*/

	return abc, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
