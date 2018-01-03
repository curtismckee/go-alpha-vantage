package timeseries // alpha vantage

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type intradayConfig struct {
	outputSize string
	dataType   string
	interval   string
}

type intradayOption func(opt *intradayConfig) error

func SetIntradayOutputSize(s string) intradayOption {

	return func(config *intradayConfig) error {
		config.outputSize = s
		return nil
	}
}

func SetIntradayDataType(s string) intradayOption {

	return func(config *intradayConfig) error {
		config.dataType = s
		return nil
	}
}

func SetIntradayInterval(s string) intradayOption {

	return func(config *intradayConfig) error {
		config.interval = s
		return nil
	}
}

func Intraday(symbol string, apikey string, opts ...intradayOption) (*http.Response, error) {

	defaultOptions := &intradayConfig{
		outputSize: "compact",
		dataType:   "json",
		interval:   "15min",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=TIME_SERIES_INTRADAY&symbol=%s&apikey=%s&interval=%s&datatype=%s&outputsize=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.interval,
		defaultOptions.dataType,
		defaultOptions.outputSize)

	resp, err := http.Get(url)
	return resp, err
}
