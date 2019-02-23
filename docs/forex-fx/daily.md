# Daily

{% api-method method="get" host="https://alphavantage.co" path="/query?" %}
{% api-method-summary %}
Forex (FX) Daily
{% endapi-method-summary %}

{% api-method-description %}
s API returns the daily time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-query-parameters %}
{% api-method-parameter name="function" type="string" required=true %}
The time series of your choice. In this case, `function=TIME_SERIES_INTRADAY`
{% endapi-method-parameter %}

{% api-method-parameter name="from\_symbol" type="string" required=true %}
 three-letter symbol from the forex currency list. For example: `from_symbol=EUR`
{% endapi-method-parameter %}

{% api-method-parameter name="to\_symbol" type="string" required=true %}
 three-letter symbol from the forex currency list. For example: `from_symbol=USD`
{% endapi-method-parameter %}

{% api-method-parameter name="outputsize" type="string" %}
By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points in the daily time series; `full` returns the full-length time series of 20+ years of historical data. The "compact" option is recommended if you would like to reduce the data size of each API call. 
{% endapi-method-parameter %}

{% api-method-parameter name="datatype" type="string" %}
By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. 
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
    "1. Information": "Forex Daily Prices (open, high, low, close)",
    "2. From Symbol": "EUR",
    "3. To Symbol": "USD",
    "4. Output Size": "Compact",
    "5. Last Refreshed": "2019-02-22 01:55:00",
    "6. Time Zone": "GMT+8"
  },
  "Time Series FX (Daily)": {
    "2019-02-22": {
      "1. open": "1.1347",
      "2. high": "1.1348",
      "3. low": "1.1321",
      "4. close": "1.1329"
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
Invalid API Call Error Message. Note: The http status code actually returns 200 with the error message.
{% endapi-method-response-example-description %}

```javascript
{
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for FX_DAILY."
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



