<center>
  <h1>Technical Indicators - Weighted Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the weighted moving average (WMA) values. See also: [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=WeightedMA.htm).


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=WMA` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each moving average value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=WMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=WMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Weighted Moving Average (WMA)",
    "3: Last Refreshed": "2019-03-11",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern"
  },
  "Technical Analysis: WMA": {
    "2019-03-11": {
      "WMA": "108.6904"
    },
    "2019-03-08": {
      "WMA": "107.7736"
    },
    "2019-03-01": {
      "WMA": "106.2089"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
