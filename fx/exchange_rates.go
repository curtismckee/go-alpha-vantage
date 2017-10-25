package fx // alpha vantage

import (
	"fmt"
	"net/http"

	"../av"
)

type exchangeRatesConfig struct {
	fromCurrency string
	toCurrency   string
}

type exchangeRatesOption func(opt *exchangeRatesConfig) error

func SetFromCurrency(s string) exchangeRatesOption {

	return func(config *exchangeRatesConfig) error {
		config.fromCurrency = s
		return nil
	}
}

func SetToCurrency(s string) exchangeRatesOption {

	return func(config *exchangeRatesConfig) error {
		config.toCurrency = s
		return nil
	}
}

func ExchangeRates(symbol string, apikey string, opts ...exchangeRatesOption) (*http.Response, error) {

	defaultOptions := &exchangeRatesConfig{
		fromCurrency: "USD",
		toCurrency:   "EUR",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=CURRENCY_EXCHANGE_RATE&symbol=%s&apikey=%s&from_currency=%s&to_currency=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.fromCurrency,
		defaultOptions.toCurrency)

	resp, err := http.Get(url)
	return resp, err
}
