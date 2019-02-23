# Daily

{% api-method method="get" host="https://alphavantage.co" path="/query?" %}
{% api-method-summary %}
Time-Series Daily
{% endapi-method-summary %}

{% api-method-description %}
This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume) of the global equity specified, covering 20+ years of historical data.
The most recent data point is the cumulative prices and volume information of the current trading day, updated realtime.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-query-parameters %}
{% api-method-parameter name="function" type="string" required=true %}
The time series of your choice. In this case, `function=TIME_SERIES_INTRADAY`
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
The name of the equity of your choice. For example: `symbol=MSFT`
{% endapi-method-parameter %}

{% api-method-parameter name="outputsize" type="string" %}
By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points; `full` returns the full-length time series of 20+ years of historical data. The "compact" option is recommended if you would like to reduce the data size of each API call. 
{% endapi-method-parameter %}

{% api-method-parameter name="datatype" type="string" %}
By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. 
{% endapi-method-parameter %}

{% api-method-parameter name="apikey" type="string" required=true %}
Your API key.
{% endapi-method-parameter %}

{% endapi-method-query-parameters %}
{% endapi-method-request %}


{% api-method-response %}

{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Success
{% endapi-method-response-example-description %}

```javascript
{
  "Meta Data": {
    "1. Information": "Daily Prices (open, high, low, close) and Volume",
    "2. Symbol": "MSFT",
    "3. Last Refreshed": "2019-02-20 13:53:08",
    "4. Output Size": "Compact",
    "5. Time Zone": "US/Eastern"
    },
    "Time Series (Daily)": {
      "2019-02-20": {
      "1. open": "107.8600",
      "2. high": "107.9400",
      "3. low": "106.2950",
      "4. close": "106.8600",
      "5. volume": "11090306"
    },
    { ... },
    { ... },
    { ... },
  }
}
```
{% endapi-method-response-example %}

{% api-method-response-example httpCode=404 %}
{% api-method-response-example-description %}
Invalid API Call Error Message. Note: The http status code actually returns 200 with the error message.
{% endapi-method-response-example-description %}

```javascript
{
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for TIME_SERIES_INTRADAY."
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



