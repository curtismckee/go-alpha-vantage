package indicators // alpha vantage technical indicators

import (
	"fmt"
	"net/http"

	"../av"
)

type stochfConfig struct {
	fastKPeriod string
	fastDPeriod string
	fastDMAType string
}

type stochfOption func(opt *stochfConfig) error

func SetSTOCHFFastKPeriod(s string) stochfOption {

	return func(config *stochfConfig) error {
		config.fastKPeriod = s
		return nil
	}
}

func SetSTOCHFFastDPeriod(s string) stochfOption {

	return func(config *stochfConfig) error {
		config.fastDPeriod = s
		return nil
	}
}

func SetSTOCHFFastDMAType(s string) stochfOption {

	return func(config *stochfConfig) error {
		config.fastDMAType = s
		return nil
	}
}

func STOCHF(symbol string, apikey string, interval string, opts ...stochfOption) (*http.Response, error) {

	defaultOptions := &stochfConfig{
		fastKPeriod: "5",
		fastDPeriod: "3",
		fastDMAType: "0",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=STOCHF&symbol=%s&apikey=%s&interval=%s&fastkperiod=%s&fastdperiod=%s&fastdmatype=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		defaultOptions.fastKPeriod,
		defaultOptions.fastDPeriod,
		defaultOptions.fastDMAType)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
