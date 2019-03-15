<center>
  <h1>Forex - Weekly</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the weekly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.  
The latest data point is the cumulative price information for the week (or partial week) containing the current trading day, updated realtime. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=FX_WEEKLY` |
| to\_symbol      | string  | true      | A three-letter symbol from the forex currency list. For example: `from_symbol=EUR`. |
| from\_symbol    | string  | true      | A three-letter symbol from the forex currency list. For example: `from_symbol=USD`. |
| datatype        | string  | true      | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the weekly time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=FX_WEEKLY&from_symbol=EUR&to_symbol=USD&apikey=demo](https://www.alphavantage.co/query?function=FX_WEEKLY&from_symbol=EUR&to_symbol=USD&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1. Information": "Forex Weekly Prices (open, high, low, close)",
    "2. From Symbol": "EUR",
    "3. To Symbol": "USD",
    "4. Last Refreshed": "2019-02-22 02:45:00",
    "5. Time Zone": "GMT+8"
  },
  "Time Series FX (Weekly)": {
    "2019-02-22": {
      "1. open": "1.1291",
      "2. high": "1.1371",
      "3. low": "1.1274",
      "4. close": "1.1330"
    },
    { ... },
    { ... },
    { ... }
  }
}
```

<!-- tabs:end -->
