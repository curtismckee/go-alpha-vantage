<center>
  <h1>Technical Indicator - Hillbert Transform, Dominant Cycle Period</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Hilbert transform, dominant cycle period (HT_DCPERIOD) values. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=HT_DCPERIOD` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series_type     | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low`
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=HT_DCPERIOD&symbol=MSFT&interval=daily&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=HT_DCPERIOD&symbol=MSFT&interval=daily&series_type=close&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Hilbert Transform - Dominant Cycle Period (HT_DCPERIOD)",
    "3: Last Refreshed": "2019-03-14 16:00:02",
    "4: Interval": "daily",
    "5: Series Type": "close",
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: HT_DCPERIOD": {
    "2019-03-14 16:00:02": {
      "DCPERIOD": "25.0331"
    },
    "2019-03-13": {
      "DCPERIOD": "24.9081"
    },
    "2019-03-12": {
      "DCPERIOD": "24.7536"
    },
    "2019-03-11": {
      "DCPERIOD": "24.4673"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
