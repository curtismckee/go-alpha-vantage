<center>
  <h1>Technical Indicators - Simple Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the simple moving average (SMA) values. See also: [Investopedia article](http://www.investopedia.com/articles/technical/052201.asp) and [mathematical reference](http://www.fmlabs.com/reference/default.htm?url=SimpleMA.htm).

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=SMA` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each moving average value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily times series in JSON format; `csv` returns the time series as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Equity Endpoint:  

[https://www.alphavantage.co/query?function=SMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=SMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Simple Moving Average (SMA)",
    "3: Last Refreshed": "2019-03-07",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern"
  },
  "Technical Analysis: SMA": {
    "2019-03-07": {
      "SMA": "105.9480"
    },
    "2019-03-01": {
      "SMA": "104.4140"
    },
    "2019-02-22": {
      "SMA": "103.7790"
    },
    "2019-02-15": {
      "SMA": "103.4800"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
