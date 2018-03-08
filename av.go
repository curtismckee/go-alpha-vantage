package av

import (
	"net/http"
	"net/url"
	"time"
)

const (
	// HostDefault is the default host for Alpha Vantage
	HostDefault = "www.alphavantage.co"
)

const (
	schemeHttps = "https"

	queryApiKey     = "apikey"
	queryDataType   = "datatype"
	queryOutputSize = "outputsize"
	querySymbol     = "symbol"
	queryMarket     = "market"
	queryEndpoint   = "function"
	queryInterval   = "interval"

	valueCompact                  = "compact"
	valueJson                     = "csv"
	valueDigitcalCurrencyEndpoint = "DIGITAL_CURRENCY_INTRADAY"

	pathQuery = "query"

	requestTimeout = time.Second * 30
)

// Connection is an interface that requests data from a server
type Connection interface {
	// Request creates an http Response from the given endpoint URL
	Request(endpoint *url.URL) (*http.Response, error)
}

type avConnection struct {
	client *http.Client
	host   string
}

// NewConnectionHost creates a new connection at the default Alpha Vantage host
func NewConnection() Connection {
	return NewConnectionHost(HostDefault)
}

// NewConnectionHost creates a new connection at the given host
func NewConnectionHost(host string) Connection {
	client := &http.Client{
		Timeout: requestTimeout,
	}
	return &avConnection{
		client: client,
		host:   host,
	}
}

// Request will make an HTTP GET request for the given endpoint from Alpha Vantage
func (conn *avConnection) Request(endpoint *url.URL) (*http.Response, error) {
	endpoint.Scheme = schemeHttps
	endpoint.Host = conn.host
	targetUrl := endpoint.String()
	return conn.client.Get(targetUrl)
}

// Client is a service used to query Alpha Vantage stock data
type Client struct {
	conn   Connection
	apiKey string
}

// NewClientConnection creates a new Client with the default Alpha Vantage connection
func NewClient(apiKey string) *Client {
	return NewClientConnection(apiKey, NewConnection())
}

// NewClientConnection creates a Client with a specific connection
func NewClientConnection(apiKey string, connection Connection) *Client {
	return &Client{
		conn:   connection,
		apiKey: apiKey,
	}
}

// buildRequestPath builds an endpoint URL with the given query parameters
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

// StockTimeSeriesIntraday queries a stock symbols statistics throughout the day.
// Data is returned from past to present.
func (c *Client) StockTimeSeriesIntraday(timeInterval TimeInterval, symbol string) ([]*TimeSeriesValue, error) {
	endpoint := c.buildRequestPath(map[string]string{
		queryEndpoint: timeSeriesIntraday.keyName(),
		queryInterval: timeInterval.keyName(),
		querySymbol:   symbol,
	})
	response, err := c.conn.Request(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return parseTimeSeriesData(response.Body)
}

// StockTimeSeries queries a stock symbols statistics for a given time frame.
// Data is returned from past to present.
func (c *Client) StockTimeSeries(timeSeries TimeSeries, symbol string) ([]*TimeSeriesValue, error) {
	endpoint := c.buildRequestPath(map[string]string{
		queryEndpoint: timeSeries.keyName(),
		querySymbol:   symbol,
	})
	response, err := c.conn.Request(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return parseTimeSeriesData(response.Body)
}

// DigitalCurrency queries statistics of a digital currency in terms of a physical currency throughout the day.
// Data is returned from past to present.
func (c *Client) DigitalCurrency(digital, physical string) ([]*DigitalCurrencySeriesValue, error) {
	endpoint := c.buildRequestPath(map[string]string{
		queryEndpoint: valueDigitcalCurrencyEndpoint,
		querySymbol:   digital,
		queryMarket:   physical,
	})
	response, err := c.conn.Request(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return parseDigitalCurrencySeriesData(response.Body)
}
