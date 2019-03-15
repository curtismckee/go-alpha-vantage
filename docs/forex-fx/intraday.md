<center>
  <h1>Forex - Intraday</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**


This API returns intraday time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=FX_INTRADAY` |
| to\_symbol      | string  | true      | A three-letter symbol from the forex currency list. For example: `symbol=EUR`. |
| from\_symbol    | string  | true      | A three-letter symbol from the forex currency list. For example: `to_symbol=USD` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following values are supported: `1min`, `5min`, `15min`, `30min`, `60min` |
| outputsize      | string  | optional  | By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points in the intraday time series; `full` returns the full-length intraday time series. The "compact" option is recommended if you would like to reduce the data size of each API call.
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=USD&interval=5min&apikey=demo](https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=USD&interval=5min&apikey=demo)

Example Response:  

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
<!-- tabs:end -->
