package av // alpha vantage

import (
	"fmt"
	"net/http"
)

type dailyConfig struct {
	outputSize string
	dataType   string
}

type dailyOption func(opt *dailyConfig) error

func SetDailyOutputSize(s string) dailyOption {

	return func(config *dailyConfig) error {
		config.outputSize = s
		return nil
	}
}

func SetDailyDataType(s string) dailyOption {

	return func(config *dailyConfig) error {
		config.dataType = s
		return nil
	}
}

func Daily(symbol string, apikey string, opts ...dailyOption) (*http.Response, error) {

	defaultOptions := &dailyConfig{
		outputSize: "compact",
		dataType:   "json",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s&datatype=%s&outputsize=%s",
		AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.dataType,
		defaultOptions.outputSize)

	resp, err := http.Get(url)
	return resp, err
}
