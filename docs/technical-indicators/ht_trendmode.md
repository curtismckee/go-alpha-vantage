<center>
  <h1>Technical Indicator - Hillbert Transform, Trend Vs Cycle Mode</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Hilbert transform, trend vs cycle mode (HT_TRENDMODE) values. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=HT_TRENDMODE` |
| symbol          | string  | true      | The name of the technical indicator of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series_type     | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=HT_TRENDMODE&symbol=MSFT&interval=weekly&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=HT_TRENDMODE&symbol=MSFT&interval=weekly&series_type=close&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Hilbert Transform - Trend vs Cycle Mode (HT_TRENDMODE)",
    "3: Last Refreshed": "2019-03-14 14:18:49",
    "4: Interval": "weekly",
    "5: Series Type": "close",
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: HT_TRENDMODE": {
    "2019-03-14 14:18:49": {
      "TRENDMODE": "1"
    },
    "2019-03-08": {
      "TRENDMODE": "1"
    },
    "2019-03-01": {
      "TRENDMODE": "1"
    },
    "2019-02-22": {
      "TRENDMODE": "0"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
