package crypto // alpha vantage

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type intradayConfig struct {
	market string
}

type intradayOption func(opt *intradayConfig) error

func SetIntradayMarket(s string) intradayOption {

	return func(config *intradayConfig) error {
		config.market = s
		return nil
	}
}

func Intraday(symbol string, apikey string, opts ...intradayOption) (*http.Response, error) {

	defaultOptions := &intradayConfig{
		market: "USD",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=DIGITAL_CURRENCY_INTRADAY&symbol=%s&apikey=%s&market=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.market)

	resp, err := http.Get(url)
	return resp, err
}
