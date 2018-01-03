package indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

type stochRsiConfig struct {
	fastKPeriod string
	fastDPeriod string
	fastDMAType string
}

type stochRsiOption func(opt *stochRsiConfig) error

func SetSTOCHRSIFastKPeriod(s string) stochRsiOption {

	return func(config *stochRsiConfig) error {
		config.fastKPeriod = s
		return nil
	}
}

func SetSTOCHRSIFastDPeriod(s string) stochRsiOption {

	return func(config *stochRsiConfig) error {
		config.fastDPeriod = s
		return nil
	}
}

func SetSTOCHRSIFastDMAType(s string) stochRsiOption {

	return func(config *stochRsiConfig) error {
		config.fastDMAType = s
		return nil
	}
}

func StochRsi(symbol, apikey, interval, timeperiod, seriestype string, opts ...stochRsiOption) (*http.Response, error) {

	defaultOptions := &stochRsiConfig{
		fastKPeriod: "5",
		fastDPeriod: "3",
		fastDMAType: "0",
	}

	for _, opt := range opts {
		opt(defaultOptions)
	}

	url := fmt.Sprintf("%s/query?function=STOCHRSI&symbol=%s&apikey=%s&interval=%s&time_period=%s&series_type=%s&fastkperiod=%s&fastdperiod=%s&fastdmatype=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		timeperiod,
		seriestype,
		defaultOptions.fastKPeriod,
		defaultOptions.fastDPeriod,
		defaultOptions.fastDMAType)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
