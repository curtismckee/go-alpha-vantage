package fx

import (
	"net/http"
)

type FxConnection interface {
	ExchangeRates(opts ...exchangeRatesOption) (*http.Response, error)
}

type fxClient struct {
	apikey string
}

func NewClient(apikey string) FxConnection {
	return &fxClient{
		apikey: apikey,
	}
}
