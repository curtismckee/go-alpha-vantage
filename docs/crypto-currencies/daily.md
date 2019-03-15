<center>
  <h1>Crypto Currencies - Daily</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the daily historical time series for a digital currency (e.g., BTC) traded on a specific market (e.g., CNY/Chinese Yuan), refreshed daily at midnight (UTC).  
Prices and volumes are quoted in both the market-specific currency and USD.  

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=DIGITAL_CURRENCY_DAILY` |
| symbol          | string  | true      | The digital/crypto currency of your choice. It can be any of the currencies in the digital currency list. For example: `symbol=BTC`. |
| market          | string  | true      | The exchange market of your choice. It can be any of the market in the market list. For example: `market=CNY`. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=BTC&market=CNY&apikey=demo](https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=BTC&market=CNY&apikey=demo)

Example Response:  

```javascript
{
  "MetaData": {
    "1. Information": "Daily Prices and Volumes for Digital Currency",
    "2. Digital Currency Code": "BTC",
    "3. Digital Currency Name": "Bitcoin",
    "4. Market Code": "CNY",
    "5. Market Name": "Chinese Yuan",
    "6. Last Refreshed": "2019-02-19 (end of day)",
    "7. Time Zone": "UTC"
  },
  "Time Series (Digital Currency Daily)": {
    "2019-02-19": {
      "1a. open (CNY)": "29417.89447409",
      "1b. open (USD)": "4347.51500595",
      "2a. high (CNY)": "29907.45451834",
      "2b. high (USD)": "4425.22719754",
      "3a. low (CNY)": "29301.33155268",
      "3b. low (USD)": "4325.16735438",
      "4a. close (CNY)": "29572.85596385",
      "4b. close (USD)": "4375.58911908",
      "5. volume": "5372.38336854",
      "6. market cap (USD)": "23507342.21089026"
    },
    { ... },
    { ... },
    { ... }
  }
}
```

<!-- tabs:end -->
