# Monthly Adjusted

{% api-method method="get" host="https://alphavantage.co" path="/query?" %}
{% api-method-summary %}
Time-Series Monthly Adjusted
{% endapi-method-summary %}

{% api-method-description %}
This API returns monthly adjusted time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly adjusted close, monthly volume, monthly dividend) of the equity specified, covering 20+ years of historical data. 
The latest data point is the cumulative prices and volume information for the month (or partial month) that contains the current trading day, updated realtime. 
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-query-parameters %}
{% api-method-parameter name="function" type="string" required=true %}
The time series of your choice. In this case, `function=TIME_SERIES_MONTHLY_ADJUSTED`
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
The name of the equity of your choice. For example: `symbol=MSFT`
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
    "1. Information": "Monthly Adjusted Prices and Volumes",
    "2. Symbol": "MSFT",
    "3. Last Refreshed": "2019-02-20 14:27:22",
    "4. Time Zone": "US/Eastern"
  },
  "Monthly Adjusted Time Series": {
    "2019-02-20": {
      "1. open": "103.7750",
      "2. high": "108.6600",
      "3. low": "102.3500",
      "4. close": "106.5100",
      "5. adjusted close": "106.5100",
      "6. volume": "307629143",
      "7. dividend amount": "0.0000"
    },
    { ... }.
    { ... },
    { ... },
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
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for TIME_SERIES_INTRADAY."
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



