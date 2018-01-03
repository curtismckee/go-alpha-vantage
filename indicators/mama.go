package indicators // alpha vantage technical indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type mamaConfig struct {
	fastLimit string
	slowLimit string
}

type mamaOption func(opt *mamaConfig) error

func SetMAMAFastLimit(s string) mamaOption {

	return func(config *mamaConfig) error {
		config.fastLimit = s
		return nil
	}
}

func SetMAMASlowLimit(s string) mamaOption {

	return func(config *mamaConfig) error {
		config.slowLimit = s
		return nil
	}
}

func MAMA(symbol string, apikey string, interval string, seriesType string, opts ...mamaOption) (*http.Response, error) {

	defaultOptions := &mamaConfig{
		fastLimit: "0.01",
		slowLimit: "0.01",
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
		defaultOptions.fastLimit,
		defaultOptions.slowLimit)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
