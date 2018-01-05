package av

import (
	"net/url"
	"net/http"
	"time"
)

const (
	HostLive = "www.alphavantage.co"
)

const (
	schemeHttps = "https"

	queryApiKey     = "apikey"
	queryDataType   = "datatype"
	queryOutputSize = "outputsize"
	querySymbol     = "symbol"
	queryEndpoint   = "function"

	valueCompact = "compact"
	valueJson    = "json"

	pathQuery = "query"

	requestTimeout = time.Second * 30
)

type Connection interface {
	Request(endpoint *url.URL) (*http.Response, error)
}

type avConnection struct {
	client *http.Client
	host   string
}

func NewConnection() Connection {
	return NewConnectionHost(HostLive)
}

func NewConnectionHost(host string) Connection {
	client := &http.Client{
		Timeout: requestTimeout,
	}
	return &avConnection{
		client: client,
		host:   host,
	}
}

func (conn *avConnection) Request(endpoint *url.URL) (*http.Response, error) {
	endpoint.Scheme = schemeHttps
	endpoint.Host = conn.host
	return conn.client.Get(endpoint.String())
}

type Client struct {
	conn   Connection
	apiKey string
}

func NewClient(apiKey string) *Client {
	return NewClientConnection(apiKey, NewConnection())
}

func NewClientConnection(apiKey string, connection Connection) *Client {
	return &Client{
		conn:   connection,
		apiKey: apiKey,
	}
}

func (c *Client) buildRequestPath(params map[string]string) *url.URL {
	// build our URL
	endpoint := &url.URL{}
	endpoint.Path = pathQuery

	// base parameters
	endpoint.Query()
	query := endpoint.Query()
	query.Set(queryApiKey, c.apiKey)
	query.Set(queryDataType, valueCompact)
	query.Set(queryOutputSize, valueJson)

	// additional parameters
	for key, value := range params {
		query.Set(key, value)
	}

	return endpoint
}

func (c *Client) StockTimeSeries(timeSeries TimeSeries, symbol string) (*TimeSeriesData, error) {
	endpoint := c.buildRequestPath(map[string]string{
		queryEndpoint: timeSeries.String(),
		querySymbol:   symbol,
	})

	response, err := c.conn.Request(endpoint)
	if err != nil {
		return nil, err
	}

	return parseTimeSeriesData(response)
}
