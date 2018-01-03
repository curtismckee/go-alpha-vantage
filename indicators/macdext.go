package indicators // alpha vantage technical indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type macdextConfig struct {
	fastPeriod   string
	slowPeriod   string
	signalPeriod string
	fastMAType   string
	slowMAType   string
	signalMAType string
}

type macdextOption func(opt *macdextConfig) error

func SetMACDEXTFastPeriod(s string) macdextOption {

	return func(config *macdextConfig) error {
		config.fastPeriod = s
		return nil
	}
}

func SetMACDEXTSlowPeriod(s string) macdextOption {

	return func(config *macdextConfig) error {
		config.slowPeriod = s
		return nil
	}
}

func SetMACDEXTSignalPeriod(s string) macdextOption {

	return func(config *macdextConfig) error {
		config.signalPeriod = s
		return nil
	}
}

func SetMACDEXTFastMAType(s string) macdextOption {

	return func(config *macdextConfig) error {
		config.fastMAType = s
		return nil
	}
}

func SetMACDEXTSlowMAType(s string) macdextOption {

	return func(config *macdextConfig) error {
		config.slowMAType = s
		return nil
	}
}

func SetMACDEXTSignalMAType(s string) macdextOption {

	return func(config *macdextConfig) error {
		config.signalMAType = s
		return nil
	}
}

func MACDEXT(symbol string, apikey string, interval string, seriesType string, opts ...macdextOption) (*http.Response, error) {

	defaultOptions := &macdextConfig{
		fastPeriod:   "12",
		slowPeriod:   "26",
		signalPeriod: "9",
		fastMAType:   "0",
		slowMAType:   "0",
		signalMAType: "0",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=MAMA&symbol=%s&apikey=%s&interval=%s&series_type=%s&fastperiod=%sslowperiod=%s&signalperiod=%s&fastmatype=%s&slowmatype=%s&signalmatype=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		seriesType,
		defaultOptions.fastPeriod,
		defaultOptions.slowPeriod,
		defaultOptions.signalPeriod,
		defaultOptions.fastMAType,
		defaultOptions.slowMAType,
		defaultOptions.signalMAType)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
