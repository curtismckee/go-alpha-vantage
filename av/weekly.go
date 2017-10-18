package av // alpha vantage

import (
	"fmt"
	"net/http"
)

type weeklyConfig struct {
	dataType string
}

type weeklyOption func(opt *weeklyConfig) error

func SetWeeklyDataType(s string) weeklyOption {

	return func(config *weeklyConfig) error {
		config.dataType = s
		return nil
	}
}

func Weekly(symbol string, apikey string, opts ...weeklyOption) (*http.Response, error) {

	defaultOptions := &weeklyConfig{
		dataType: "json",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=%s&datatype=%s&outputsize=%s",
		AV_BASE_URL,
		symbol,
		apikey,
		defaultOptions.dataType)

	resp, err := http.Get(url)
	return resp, err
}
