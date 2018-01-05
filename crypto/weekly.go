package crypto // alpha vantage

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type weeklyConfig struct {
	market string
}

type weeklyOption func(opt *weeklyConfig) error

func SetWeeklyMarket(s string) weeklyOption {

	return func(config *weeklyConfig) error {
		config.market = s
		return nil
	}
}

func (c *cryptoClient) Weekly(symbol string, opts ...weeklyOption) (*http.Response, error) {

	defaultOptions := &weeklyConfig{
		market: "USD",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=DIGITAL_CURRENCY_WEEKLY&symbol=%s&apikey=%s&market=%s",
		av.AV_BASE_URL,
		symbol,
		c.apikey,
		defaultOptions.market)

	resp, err := http.Get(url)
	return resp, err
}
