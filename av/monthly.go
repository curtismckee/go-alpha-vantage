package av // alpha vantage

import (
	"fmt"
	"net/http"
)

type monthlyConfig struct {
	dataType string
}

type monthlyOption func(opt *monthlyConfig) error

func SetMonthlyDataType(s string) monthlyOption {

	return func(config *monthlyConfig) error {
		config.dataType = s
		return nil
	}
}

func Monthly(symbol string, apikey string, opts ...monthlyOption) (*http.Response, error) {

	defaultOptions := &monthlyConfig{
		dataType: "json",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=TIME_SERIES_MONTHLY&symbol=%s&apikey=%s&datatype=%s&outputsize=%s",
		AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.dataType)

	resp, err := http.Get(url)
	return resp, err
}
