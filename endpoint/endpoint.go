package endpoint

import (
	"ModularMicroservice/service"
	"ModularMicroservice/transport"
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*
// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
// the endpoint will receive a request, convert to the desired
// format, invoke the service and return the response structure
*/
func MakeFetchStockEndpoint(svc service.StockService) endpoint.Endpoint {

	/*
		interface{} -> It means it is am empty interface it accepts any type of value not only interface.
		See page 580 in the text book
		request interface{} means request is the variable name.


		Don’t try calling any methods on an empty-interface value! Remember,
		if you have a value with an interface type, you can only call methods on it
		that are part of the interface. And the empty interface doesn’t have any
		methods. That means there are no methods you can call on a value with the
		empty interface type!

		To call methods on a value with the empty interface type, you’d need to use
		a type assertion to get a value of the concrete type back.
		The syntax for the above is : val.(Type) => res := val.(Type)
		Now we can use response to get the values or functions present inside the "Type"

		req := request.(uppercaseRequest) -> This is call Type assertion.
		Since uppercaseRequest is a struct so after type assertion we can use req to get the methods and variables present inside "uppercaseRequest" Type

		The role of the endpoint is to receive a request, convert it to the expected struct, invoke the service layer, and return another struct.
		The endpoint layer does not know anything about the upper layer, because it makes no difference whether the endpoint is being invoked via HTTP, gRPC,
		or another form of transport.
	*/
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.FetchstockRequest)
		v, err := svc.FetchStocks(req.Stocks)
		var e []string
		if err != nil {
			return transport.FetchstockRequest{e}, nil
		}
		return transport.FetchstockResponse{v, ""}, nil
	}
}
