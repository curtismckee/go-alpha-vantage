package av // alpha vantage

import (
	"fmt"
	"net/http"
)

type dailyConfig struct {
	outputSize string
	dataType   string
	adjusted   bool
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

func SetDailyAdjusted(b bool) dailyOption {

	return func(config *dailyConfig) error {
		config.adjusted = b
		return nil
	}
}

func Daily(symbol string, apikey string, opts ...dailyOption) (*http.Response, error) {

	var adjustedString string

	defaultOptions := &dailyConfig{
		outputSize: "compact",
		dataType:   "json",
		adjusted:   false,
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	if defaultOptions.adjusted == true {
		adjustedString = "TIME_SERIES_DAILY_ADJUSTED"
	} else {
		adjustedString = "TIME_SERIES_DAILY"
	}

	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s&datatype=%s&outputsize=%s",
		AV_BASE_URL,
		adjustedString,
		symbol,
		apikey,
		defaultOptions.dataType,
		defaultOptions.outputSize)

	resp, err := http.Get(url)
	return resp, err
}
