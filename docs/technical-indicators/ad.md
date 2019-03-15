<center>
  <h1>Technical Indicator - Chalkin A/D Line</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Chaikin A/D line (AD) values. See also: [Investopedia article](https://www.investopedia.com/articles/active-trading/031914/understanding-chaikin-oscillator.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=AccumDist.html)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=AD` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min` |
| outputsize      | string  | optional  | By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points in the intraday time series; `full` returns the full-length intraday time series. The "compact" option is recommended if you would like to reduce the data size of each API call.
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=MSFT&interval=5min&apikey=demo](https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=MSFT&interval=5min&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1. Information": "Intraday (5min) open, high, low, close prices and volume",
    "2. Symbol": "MSFT",
    "3. Last Refreshed": "2019-02-20 12:25:00",
    "4. Interval": "5min",
    "5. Output Size": "Compact",
    "6. Time Zone": "US/Eastern"
    },
    "Time Series (5min)": {
      "2019-02-20 12:25:00": {
      "1. open": "107.0099",
      "2. high": "107.0099",
      "3. low": "106.7700",
      "4. close": "106.8300",
      "5. volume": "207752"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
