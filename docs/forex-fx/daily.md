<center>
  <h1>Forex - Daily</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the daily time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=TIME_SERIES_INTRADAY` |
| to\_symbol      | string  | true      | The three-letter symbol from the forex currency list. For example: `from_symbol=EUR`. |
| from\_symbol    | string  | true      | The three-letter symbol from the forex currency list. For example: `from_symbol=USD`. |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min` |
| outputsize      | string  | optional  | By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points in the daily time series; `full` returns the full-length time series of 20+ years of historical data. The "compact" option is recommended if you would like to reduce the data size of each API call. |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=EUR&to_symbol=USD&apikey=demo](https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=EUR&to_symbol=USD&apikey=demo)

Example Response:  

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

<!-- tabs:end -->
