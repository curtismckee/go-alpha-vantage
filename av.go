package av

import (
	"net/http"
	"net/url"
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
	queryInterval   = "interval"

	valueCompact = "compact"
	valueJson    = "csv"

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
	targetUrl := endpoint.String()
	return conn.client.Get(targetUrl)
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
	query := endpoint.Query()
	query.Set(queryApiKey, c.apiKey)
	query.Set(queryDataType, valueJson)
	query.Set(queryOutputSize, valueCompact)

	// additional parameters
	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint
}

func (c *Client) StockTimeSeriesIntraday(timeInterval TimeInterval, symbol string) ([]*TimeSeriesValue, error) {
	endpoint := c.buildRequestPath(map[string]string{
		queryEndpoint: timeSeriesIntraday.KeyName(),
		queryInterval: timeInterval.KeyName(),
		querySymbol:   symbol,
	})

	response, err := c.conn.Request(endpoint)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return parseTimeSeriesData(response.Body)
}

func (c *Client) StockTimeSeries(timeSeries TimeSeries, symbol string) ([]*TimeSeriesValue, error) {
	endpoint := c.buildRequestPath(map[string]string{
		queryEndpoint: timeSeries.KeyName(),
		querySymbol:   symbol,
	})

	response, err := c.conn.Request(endpoint)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return parseTimeSeriesData(response.Body)
}
