package crypto

import (
	"net/http"
)

type CryptoConnection interface {
	Daily(symbol string, opts ...dailyOption) (*http.Response, error)
	Intraday(symbol string, opts ...intradayOption) (*http.Response, error)
	Weekly(symbol string, opts ...weeklyOption) (*http.Response, error)
	Monthly(symbol string, opts ...monthlyOption) (*http.Response, error)
}

type cryptoClient struct {
	apikey string
}

func NewClient(apikey string) CryptoConnection {
	return &cryptoClient{
		apikey: apikey,
	}
}
