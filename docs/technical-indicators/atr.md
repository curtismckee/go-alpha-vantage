<center>
  <h1>Technical Indicator - Average True Range</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the average true range (ATR) values. See also: [Investopedia article](https://www.investopedia.com/articles/trading/08/average-true-range.asp) and [mathematical reference](https://www.investopedia.com/articles/trading/08/average-true-range.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=ATR` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, i`daily`, `weekly`, `monthly` |
| time_period     | string  | true      | Number of data points used to calculate each ATR value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=ATR&symbol=MSFT&interval=daily&time_period=14&apikey=demo](https://www.alphavantage.co/query?function=ATR&symbol=MSFT&interval=daily&time_period=14&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Average True Range (ATR)",
    "3: Last Refreshed": "2019-03-14 12:19:53",
    "4: Interval": "daily",
    "5: Time Period": 14,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: ATR": {
    "2019-03-14 12:19:53": {
      "ATR": "1.7350"
    },
    "2019-03-13": {
      "ATR": "1.8016"
    },
    "2019-03-12": {
      "ATR": "1.8340"
    },
    "2019-03-11": {
      "ATR": "1.8720"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
