package crypto // alpha vantage

import (
	"fmt"
	"net/http"

	"../av"
)

type dailyConfig struct {
	market string
}

type dailyOption func(opt *dailyConfig) error

func SetDailyMarket(s string) dailyOption {

	return func(config *dailyConfig) error {
		config.market = s
		return nil
	}
}

func Daily(symbol string, apikey string, opts ...dailyOption) (*http.Response, error) {

	defaultOptions := &dailyConfig{
		market: "USD",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=DIGITAL_CURRENCY_DAILY&symbol=%s&apikey=%s&market=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.market)

	resp, err := http.Get(url)
	return resp, err
}
