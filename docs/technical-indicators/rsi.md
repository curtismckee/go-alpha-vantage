<center>
  <h1>Technical Indicators - Relative Strength Index</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the relative strength index (RSI) values. See also: [Investopedia article](https://www.investopedia.com/articles/active-trading/042114/overbought-or-oversold-use-relative-strength-index-find-out.asp) and [mathematical reference](https://www.investopedia.com/articles/active-trading/042114/overbought-or-oversold-use-relative-strength-index-find-out.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=RSI` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true  | Number of data points used to calculate each RSI value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200)`. |
| series\_type    | string  | true | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=RSI&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=RSI&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Relative Strength Index (RSI)",
    "3: Last Refreshed": "2019-03-12",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: RSI": {
    "2019-03-12": {
      "RSI": "57.8730"
    },
    "2019-03-08": {
      "RSI": "62.6690"
    },
    "2019-03-01": {
      "RSI": "60.8569"
    },
    "2019-02-22": {
      "RSI": "54.6090"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
