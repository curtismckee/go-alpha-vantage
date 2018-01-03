package indicators // alpha vantage technical indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type stochConfig struct {
	fastKPeriod string
	slowKPeriod string
	slowDPeriod string
	slowKMAType string
	slowDMAType string
}

type stochOption func(opt *stochConfig) error

func SetSTOCHFastKPeriod(s string) stochOption {

	return func(config *stochConfig) error {
		config.fastKPeriod = s
		return nil
	}
}

func SetSTOCHSlowKPeriod(s string) stochOption {

	return func(config *stochConfig) error {
		config.slowKPeriod = s
		return nil
	}
}

func SetSTOCHSlowDPeriod(s string) stochOption {

	return func(config *stochConfig) error {
		config.slowDPeriod = s
		return nil
	}
}

func SetSTOCHSlowKMAType(s string) stochOption {

	return func(config *stochConfig) error {
		config.slowKMAType = s
		return nil
	}
}

func SetSTOCHSlowDMAType(s string) stochOption {

	return func(config *stochConfig) error {
		config.slowDMAType = s
		return nil
	}
}

func STOCH(symbol string, apikey string, interval string, opts ...stochOption) (*http.Response, error) {

	defaultOptions := &stochConfig{
		fastKPeriod: "5",
		slowKPeriod: "3",
		slowDPeriod: "3",
		slowKMAType: "0",
		slowDMAType: "0",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=STOCH&symbol=%s&apikey=%s&interval=%s&&fastkperiod=%sslowkperiod=%s&slowdperiod=%s&slowkmatype=%s&slowdmatype=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		defaultOptions.fastKPeriod,
		defaultOptions.slowKPeriod,
		defaultOptions.slowDPeriod,
		defaultOptions.slowKMAType,
		defaultOptions.slowDMAType)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
