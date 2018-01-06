package timeseries

import (
	"net/http"
)

type TimeSeriesConnection interface {
	Daily(symbol string, opts ...dailyOption) (*http.Response, error)
	Intraday(symbol string, opts ...intradayOption) (*http.Response, error)
	Monthly(symbol string, opts ...monthlyOption) (*http.Response, error)
	Weekly(symbol string, opts ...weeklyOption) (*http.Response, error)
}

type timeSeriesClient struct {
	apikey string
}

func NewClient(apikey string) TimeSeriesConnection {
	return &timeSeriesClient{
		apikey: apikey,
	}
}
