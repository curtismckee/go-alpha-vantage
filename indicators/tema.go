package indicators // alpha vantage technical indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

func TEMA(symbol string, apikey string, interval string, timePeriod string, seriesType string) (*http.Response, error) {

	url := fmt.Sprintf("%s/query?function=TEMA&symbol=%s&apikey=%s&interval=%s&time_period=%s&series_type=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		timePeriod,
		seriesType)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}
