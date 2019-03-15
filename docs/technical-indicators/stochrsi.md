<center>
  <h1>Technical Indicator - Stochastic Relative Strength Index</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the stochastic relative strength index (STOCHRSI) values. See also: [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=StochRSI.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=STOCHRSI` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each STOCHRSI value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`). |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| fastkperiod     | string  | optional  | The time period of the fastk moving average. Positive integers are accepted. By default, `fastkperiod=5`. |  
| fastdperiod     | string  | optional  | The time period of the fastd moving average. Positive integers are accepted. By default, `fastdperiod=3`. |
| fastdmatype     | string  | optional  | Moving average type for the fastd moving average. By default, `fastdmatype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=STOCHRSI&symbol=MSFT&interval=daily&time_period=10&series_type=close&fastkperiod=6&fastdmatype=1&apikey=demo](https://www.alphavantage.co/query?function=STOCHRSI&symbol=MSFT&interval=daily&time_period=10&series_type=close&fastkperiod=6&fastdmatype=1&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Stochastic Relative Strength Index (STOCHRSI)",
    "3: Last Refreshed": "2019-03-12",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6.1: FastK Period": 6,
    "6.2: FastD Period": 3,
    "6.3: FastD MA Type": 1,
    "7: Series Type": "close",
    "8: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: STOCHRSI": {
    "2019-03-12": {
      "FastK": "100.0000",
      "FastD": "71.8483"
    },
    "2019-03-11": {
      "FastK": "82.2205",
      "FastD": "43.6966"
    },
    "2019-03-08": {
      "FastK": "4.3123",
      "FastD": "5.1727"
    },
    "2019-03-07": {
      "FastK": "0.0000",
      "FastD": "6.0330"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
