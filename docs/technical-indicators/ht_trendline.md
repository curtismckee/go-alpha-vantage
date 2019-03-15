<center>
  <h1>Technical Indicator - Hillbert Transform, Instantaneous trendline</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Hilbert transform, instantaneous trendline (HT_TRENDLINE) values.

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=HT_TRENDLINE` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series_type     | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=HT_TRENDLINE&symbol=MSFT&interval=daily&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=HT_TRENDLINE&symbol=MSFT&interval=daily&series_type=close&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Hilbert Transform - Instantaneous Trendline (HT_TRENDLINE)",
    "3: Last Refreshed": "2019-03-14 13:40:39",
    "4: Interval": "daily",
    "5: Series Type": "close",
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: HT_TRENDLINE": {
    "2019-03-14 13:40:39": {
      "HT_TRENDLINE": "109.8369"
    },
    "2019-03-13": {
      "HT_TRENDLINE": "109.5435"
    },
    "2019-03-12": {
      "HT_TRENDLINE": "109.3102"
    },
    "2019-03-11": {
      "HT_TRENDLINE": "109.1338"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
