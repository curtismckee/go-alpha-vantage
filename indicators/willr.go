package indicators

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

func Willr(symbol, apikey, interval, timeperiod string) (*http.Response, error) {

	url := fmt.Sprintf("%s/query?function=WILLR&symbol=%s&apikey=%s&interval=%s&time_period=%s",
		av.AV_BASE_URL,
		symbol,
		apikey,
		interval,
		timeperiod)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp, err

}
