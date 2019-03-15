<center>
  <h1>Crypto Currencies - Weekly</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the weekly historical time series for a digital currency (e.g., BTC) traded on a specific market (e.g., CNY/Chinese Yuan), refreshed daily at midnight (UTC).  
Prices and volumes are quoted in both the market-specific currency and USD. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=DIGITAL_CURRENCY_WEEKLY` |
| symbol          | string  | true      | The digital/crypto currency of your choice. It can be any of the currencies in the digital currency list. For example: `symbol=BTC`. |
| market          | string  | true      | The exchange market of your choice. It can be any of the market in the market list. For example: `market=CNY`. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_WEEKLY&symbol=BTC&market=CNY&apikey=demo](https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_WEEKLY&symbol=BTC&market=CNY&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1. Information": "Weekly Prices and Volumes for Digital Currency",
    "2. Digital Currency Code": "BTC",
    "3. Digital Currency Name": "Bitcoin",
    "4. Market Code": "CNY",
    "5. Market Name": "Chinese Yuan",
    "6. Last Refreshed": "2019-02-19 (end of day)",
    "7. Time Zone": "UTC"
  },
  "Time Series (Digital Currency Weekly)": {
    "2019-02-19": {
      "1a. open (CNY)": "27629.65961630",
      "1b. open (USD)": "4079.32252168",
      "2a. high (CNY)": "29907.45451830",
      "2b. high (USD)": "4425.22719754",
      "3a. low (CNY)": "27629.65961630",
      "3b. low (USD)": "4079.32252168",
      "4a. close (CNY)": "29572.85596390",
      "4b. close (USD)": "4375.58911908",
      "5. volume": "10678.94855400",
      "6. market cap (USD)": "46581824.96990000"
    },
    { ... },
    { ... },
    { ... }
  }
}
```

<!-- tabs:end -->
