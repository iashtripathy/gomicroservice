package transport

import (
	"ModularMicroservice/service/serviceModel"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// For each method, we define request and response structs
type FetchstockRequest struct {
	Stocks []string `json:"stocks"`
}

type FetchstockResponse struct {
	V []serviceModel.StockStructure `json:"v"`

	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func DecodeFetchStocksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FetchstockRequest
	/*
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		//r.URL.Query().Get("s")
		//request.S = query_params['s']
	*/
	str := r.URL.Query().Get("stocks")

	strArray := strings.Split(str, ",")
	request.Stocks = append(request.Stocks, strArray...)
	fmt.Println(str)

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
