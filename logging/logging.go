package logging

import (
	"ModularMicroservice/service"
	"ModularMicroservice/service/serviceModel"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.StockService
}

func (mw LoggingMiddleware) FetchStocks(s []string) (output []serviceModel.StockStructure, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "fetch stocks",
			"input", strings.Join(s, ","),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.FetchStocks(s)
	return
}
