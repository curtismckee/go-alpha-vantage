package indicators // alpha vantage technical indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type macdConfig struct {
	fastPeriod   string
	slowPeriod   string
	signalPeriod string
}

type macdOption func(opt *macdConfig) error

func SetMACDFastPeriod(s string) macdOption {

	return func(config *macdConfig) error {
		config.fastPeriod = s
		return nil
	}
}

func SetMACDSlowPeriod(s string) macdOption {

	return func(config *macdConfig) error {
		config.slowPeriod = s
		return nil
	}
}

func SetMACDSignalPeriod(s string) macdOption {

	return func(config *macdConfig) error {
		config.signalPeriod = s
		return nil
	}
}

func MACD(symbol string, apikey string, interval string, seriesType string, opts ...macdOption) (*http.Response, error) {

	defaultOptions := &macdConfig{
		fastPeriod:   "12",
		slowPeriod:   "26",
		signalPeriod: "9",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=MAMA&symbol=%s&apikey=%s&interval=%s&series_type=%s&fastlimit=%sslowlimit=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		seriesType,
		defaultOptions.fastPeriod,
		defaultOptions.slowPeriod)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
