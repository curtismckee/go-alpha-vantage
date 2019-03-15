<center>
  <h1>Technical Indicators - Moving Average Convergence / Divergence</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the moving average convergence / divergence (MACD) values. See also: [Investopedia article](https://www.investopedia.com/articles/forex/05/macddiverge.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=MACD.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MACD` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min` |
| outputsize      | string  | optional  | By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points in the intraday time series; `full` returns the full-length intraday time series. The "compact" option is recommended if you would like to reduce the data size of each API call.
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=MACD&symbol=MSFT&interval=daily&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=MACD&symbol=MSFT&interval=daily&series_type=open&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Moving Average Convergence/Divergence (MACD)",
    "3: Last Refreshed": "2019-03-12",
    "4: Interval": "daily",
    "5.1: Fast Period": 12,
    "5.2: Slow Period": 26,
    "5.3: Signal Period": 9,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern"
  },
  "Technical Analysis: MACD": {
    "2019-03-12": {
      "MACD": "1.5413",
      "MACD_Signal": "1.6195",
      "MACD_Hist": "-0.0782"
    },
    "2019-03-11": {
      "MACD": "1.4838",
      "MACD_Signal": "1.6390",
      "MACD_Hist": "-0.1553"
    },
    "2019-03-08": {
      "MACD": "1.5751",
      "MACD_Signal": "1.6779",
      "MACD_Hist": "-0.1028"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
