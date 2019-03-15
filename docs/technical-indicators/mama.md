<center>
  <h1>Technical Indicators - MESA Adaptive Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the MESA adaptive moving average (MAMA) values.

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MAMA` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `weekly`, `monthly` |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| fastlimit       | string  | optional  | Positive floats are accepted. By default, `fastlimit=0.01`. |
| slowlimit       | string  | optional  | Positive floats are accepted. By default, `slowlimit=0.01`. |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=MAMA&symbol=MSFT&interval=daily&series_type=close&fastlimit=0.02&apikey=demo](https://www.alphavantage.co/query?function=MAMA&symbol=MSFT&interval=daily&series_type=close&fastlimit=0.02&apikey=demo)

Example Response:  

```javascript
{
  Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "MESA Adaptive Moving Average (MAMA)",
    "3: Last Refreshed": "2019-03-12 12:40:49",
    "4: Interval": "daily",
    "5.1: Fast Limit": 0.02,
    "5.2: Slow Limit": 0.01,
    "6: Series Type": "close",
    "7: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: MAMA": {
    "2019-03-12 12:40:49": {
      "MAMA": "105.7240",
      "FAMA": "94.0432"
    },
    "2019-03-11": {
      "MAMA": "105.5571",
      "FAMA": "93.9252"
    },
    "2019-03-08": {
      "MAMA": "105.4837",
      "FAMA": "93.8668"
    },
    "2019-03-07": {
      "MAMA": "105.4329",
      "FAMA": "93.8084"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
