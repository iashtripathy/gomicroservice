package main

import (
	"net/http"
	"os"

	"ModularMicroservice/endpoint"
	"ModularMicroservice/logging"
	"ModularMicroservice/service"
	"ModularMicroservice/transport"

	"github.com/go-kit/kit/log"
	//kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {

	Logger := log.NewLogfmtLogger(os.Stderr)

	/*
		fieldKeys := []string{"method", "error"}
		requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "my_group",
			Subsystem: "stock_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys)
		requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "my_group",
			Subsystem: "stock_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys)
		countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "my_group",
			Subsystem: "stock_service",
			Name:      "count_result",
			Help:      "The result of each count method.",
		}, []string{}) // no fields here
	*/
	var Svc service.StockService
	Svc = service.StockServiceStruct{}

	Svc = logging.LoggingMiddleware{Logger, Svc}
	//svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	fetchstocksHandler := httptransport.NewServer(
		endpoint.MakeFetchStockEndpoint(Svc),
		transport.DecodeFetchStocksRequest,
		transport.EncodeResponse,
	)

	http.Handle("/fetchstocks", fetchstocksHandler)
	//http.Handle("/metrics", promhttp.Handler())
	Logger.Log("msg", "HTTP", "addr", ":8080")
	Logger.Log("err", http.ListenAndServe(":8080", nil))
}
