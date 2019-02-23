# Intraday

{% api-method method="get" host="https://alphavantage.co" path="/query?" %}
{% api-method-summary %}
Forex (FX) Intraday
{% endapi-method-summary %}

{% api-method-description %}
This API returns intraday time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime. 
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-query-parameters %}
{% api-method-parameter name="function" type="string" required=true %}
The time series of your choice. In this case, `function=FX_INTRADAY`
{% endapi-method-parameter %}

{% api-method-parameter name="from\_symbol" type="string" required=true %}
A three-letter symbol from the forex currency list. For example: `from_symbol=EUR`
{% endapi-method-parameter %}

{% api-method-parameter name="to\_symbol" type="string" required=true %}
A three-letter symbol from the forex currency list. For example: `to_symbol=USD`
{% endapi-method-parameter %}

{% api-method-parameter name="interval" type="string" required=true %}
Time interval between two consecutive data points in the time series. The following values are supported: `1min`, `5min`, `15min`, `30min`, `60min`
{% endapi-method-parameter %}

{% api-method-parameter name="outputsize" type="string" %}
By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points in the intraday time series; `full` returns the full-length intraday time series. The "compact" option is recommended if you would like to reduce the data size of each API call. 
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
    "1. Information": "FX Intraday (5min) Time Series",
    "2. From Symbol": "EUR",
    "3. To Symbol": "USD",
    "4. Last Refreshed": "2019-02-21 18:30:00",
    "5. Interval": "5min",
    "6. Output Size": "Compact",
    "7. Time Zone": "UTC"
  },
  "Time Series FX (5min)": {
    "2019-02-21 18:30:00": {
      "1. open": "1.1329",
      "2. high": "1.1330",
      "3. low": "1.1325",
      "4. close": "1.1330"
    },
    { ... },
    { ... },
    { ... }
  }
}
```
{% endapi-method-response-example %}

{% api-method-response-example httpCode=404 %}
{% api-method-response-example-description %}
Invalid API Call Error Message.
{% endapi-method-response-example-description %}

```javascript
{
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for FX_INTRADAY."
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



