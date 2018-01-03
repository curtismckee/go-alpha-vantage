package timeseries // alpha vantage

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type weeklyConfig struct {
	dataType string
	adjusted bool
}

type weeklyOption func(opt *weeklyConfig) error

func SetWeeklyDataType(s string) weeklyOption {

	return func(config *weeklyConfig) error {
		config.dataType = s
		return nil
	}
}

func SetWeeklyAdjusted(b bool) weeklyOption {

	return func(config *weeklyConfig) error {
		config.adjusted = b
		return nil
	}
}

func Weekly(symbol string, apikey string, opts ...weeklyOption) (*http.Response, error) {

	var adjustedString string

	defaultOptions := &weeklyConfig{
		dataType: "json",
		adjusted: false,
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	if defaultOptions.adjusted == true {
		adjustedString = "TIME_SERIES_WEEKLY_ADJUSTED"
	} else {
		adjustedString = "TIME_SERIES_WEEKLY"
	}

	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s&datatype=%s&outputsize=%s",
		AV_BASE_URL,
		adjustedString,
		symbol,
		apikey,
		defaultOptions.dataType)

	resp, err := http.Get(url)
	return resp, err
}
