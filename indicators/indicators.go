package indicators

import (
	"net/http"
)

type indicatorClient struct {
	apikey string
}

type IndicatorConnection interface {
	EMA(symbol, interval, timePeriod, seriesType string) (*http.Response, error)
}

func NewClient(apikey string) IndicatorConnection {
	return &indicatorClient{
		apikey: apikey,
	}
}
