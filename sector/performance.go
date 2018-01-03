package sector // alpha vantage

import (
	"fmt"
	"net/http"

	"github.com/cmckee-dev/go-alpha-vantage/av"
)

func Performance(apikey string) (*http.Response, error) {

	url := fmt.Sprintf("%s/query?function=SECTOR&apikey=%s",
		av.AV_BASE_URL,
		apikey)

	resp, err := http.Get(url)
	return resp, err
}
