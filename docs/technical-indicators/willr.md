<center>
  <h1>Technical Indicator - Volume Weighted Average Price</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Williams' %R (WILLR) values. See also: [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=WilliamsR.htm).

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=WILLR` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each WILLR value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`). |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=WILLR&symbol=MSFT&interval=daily&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=WILLR&symbol=MSFT&interval=daily&time_period=10&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Williams' %R (WILLR)",
    "3: Last Refreshed": "2019-03-12",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: WILLR": {
    "2019-03-12": {
      "WILLR": "-7.1291"
    },
    "2019-03-11": {
      "WILLR": "-9.4382"
    },
    "2019-03-08": {
      "WILLR": "-61.5730"
    },
    "2019-03-07": {
      "WILLR": "-83.3819"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
