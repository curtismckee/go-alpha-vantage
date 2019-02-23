# Cryptocurrencies Monthly

{% api-method method="get" host="https://alphavantage.co" path="/query?" %}
{% api-method-summary %}
Cryptocurrencies Monthly
{% endapi-method-summary %}

{% api-method-description %}
This API returns the monthly historical time series for a digital currency (e.g. BTC) traded on a specific market (e.g. CNY/Chinese Yuan), refreshed daily at midnight (UTC). Prices and volumes are quoted in both the market-specific currency and USD.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-query-parameters %}
{% api-method-parameter name="function" type="string" required=true %}
The time series of your choice. In this case, `function=DIGITAL_CURRENCY_MONTHLY`
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
This digital/crypto currency of your choice. It can be any of the currencies in the digital currency list. For example: `symbol=BTC`
{% endapi-method-parameter %}

{% api-method-parameter name="market" type="string" required=true %}
The exchange market of your choice. It can be any of the market in the market list. For example: `market=CNY`.
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
    "1. Information": "Monthly Prices and Volumes for Digital Currency",
    "2. Digital Currency Code": "BTC",
    "3. Digital Currency Name": "Bitcoin",
    "4. Market Code": "CNY",
    "5. Market Name": "Chinese Yuan",
    "6. Last Refreshed": "2019-02-19 (end of day)",
    "7. Time Zone": "UTC"
  },
  "Time Series (Digital Currency Monthly)": {
    "2019-02-19": {
      "1a. open (CNY)": "25617.90107000",
      "1b. open (USD)": "3823.16790555",
      "2a. high (CNY)": "29907.45451830",
      "2b. high (USD)": "4425.22719754",
      "3a. low (CNY)": "24473.52175620",
      "3b. low (USD)": "3613.34123805",
      "4a. close (CNY)": "29572.85596390",
      "4b. close (USD)": "4375.58911908",
      "5. volume": "60087.51266160",
      "6. market cap (USD)": "242892312.45300001"
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
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for DIGITAL_CURRENCY_MONTHLY."
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}
