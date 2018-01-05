package sector

import (
	"net/http"
)

type SectorConnection interface {
	Performance() (*http.Response, error)
}

type sectorClient struct {
	apikey string
}

func NewClient(apikey string) SectorConnection {
	return &sectorClient{
		apikey: apikey,
	}
}
