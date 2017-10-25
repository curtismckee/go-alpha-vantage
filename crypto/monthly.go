package crypto // alpha vantage

import (
	"fmt"
	"net/http"

	"../av"
)

type monthlyConfig struct {
	market string
}

type monthlyOption func(opt *monthlyConfig) error

func SetMonthlyMarket(s string) monthlyOption {

	return func(config *monthlyConfig) error {
		config.market = s
		return nil
	}
}

func Monthly(symbol string, apikey string, opts ...monthlyOption) (*http.Response, error) {

	defaultOptions := &monthlyConfig{
		market: "USD",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=DIGITAL_CURRENCY_MONTHLY&symbol=%s&apikey=%s&market=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.market)

	resp, err := http.Get(url)
	return resp, err
}
