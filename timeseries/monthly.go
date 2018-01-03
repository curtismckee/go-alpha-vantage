package timeseries // alpha vantage

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type monthlyConfig struct {
	dataType string
	adjusted bool
}

type monthlyOption func(opt *monthlyConfig) error

func SetMonthlyDataType(s string) monthlyOption {

	return func(config *monthlyConfig) error {
		config.dataType = s
		return nil
	}
}

func SetMonthlyAdjusted(b bool) monthlyOption {

	return func(config *monthlyConfig) error {
		config.adjusted = b
		return nil
	}
}

func Monthly(symbol string, apikey string, opts ...monthlyOption) (*http.Response, error) {

	var adjustedString string

	defaultOptions := &monthlyConfig{
		dataType: "json",
		adjusted: false,
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	if defaultOptions.adjusted == true {
		adjustedString = "TIME_SERIES_MONTHLY_ADJUSTED"
	} else {
		adjustedString = "TIME_SERIES_MONTHLY"
	}

	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s&datatype=%s&outputsize=%s",
		av.AV_BASE_URL,
		adjustedString,
		symbol,
		apikey,
		defaultOptions.dataType)

	resp, err := http.Get(url)
	return resp, err
}
