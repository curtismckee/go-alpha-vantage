# Monthly

{% api-method method="get" host="https://alphavantage.co" path="/query?" %}
{% api-method-summary %}
Forex (FX) Monthly
{% endapi-method-summary %}

{% api-method-description %}
This API returns the monthly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime. 
The latest data point is the cumulative prices information for the month (or partial month) containing the currecnt trading day, updated realtime.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-query-parameters %}
{% api-method-parameter name="function" type="string" required=true %}
The time series of your choice. In this case, `function=FX_MONTHLY`
{% endapi-method-parameter %}

{% api-method-parameter name="from\_symbol" type="string" required=true %}
A three-letter symbol from the forex currency list. For example: `from_symbol=EUR`
{% endapi-method-parameter %}

{% api-method-parameter name="to\_symbol" type="string" required=true %}
A three-letter symbol from the forex currency list. For example: `from_symbol=USD`
{% endapi-method-parameter %}

{% api-method-parameter name="datatype" type="string" %}
By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the monthly time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. 
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
    "1. Information": "Forex Monthly Prices (open, high, low, close)",
    "2. From Symbol": "EUR",
    "3. To Symbol": "USD",
    "4. Last Refreshed": "2019-02-22 02:50:00",
    "5. Time Zone": "GMT+8"
  },
  "Time Series FX (Monthly)": {
    "2019-02-22": {
      "1. open": "1.1446",
      "2. high": "1.1488",
      "3. low": "1.1233",
      "4. close": "1.1331"
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
Invalid API Call Error Message.
{% endapi-method-response-example-description %}

```javascript
{
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for FX_MONTHLY."
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



