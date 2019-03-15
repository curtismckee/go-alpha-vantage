<center>
  <h1>Technical Indicator - Parabolic SAR</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the parabolic SAR (SAR) values. See also: [Investopedia article](https://www.investopedia.com/trading/introduction-to-parabolic-sar) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=SAR.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=SAR` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| acceleration    | string  | optional  | The acceleration factor. Positive floats are accepted. By default, `acceleration=0.01`. |
| maximum         | string  | optional  | The acceleration factor maximum value. Positive floats are accepted. By default, `acceleration=0.20`. |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=SAR&symbol=MSFT&interval=weekly&acceleration=0.05&maximum=0.25&apikey=demo](https://www.alphavantage.co/query?function=SAR&symbol=MSFT&interval=weekly&acceleration=0.05&maximum=0.25&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Parabolic SAR (SAR)",
    "3: Last Refreshed": "2019-03-14 11:39:34",
    "4: Interval": "weekly",
    "5.1: Acceleration": 0.05,
    "5.2: Maximum": 0.25,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: SAR": {
    "2019-03-14 11:39:34": {
      "SAR": "105.1443"
    },
    "2019-03-08": {
      "SAR": "102.4424"
    },
    "2019-03-01": {
      "SAR": "99.7430"
    },
    "2019-02-22": {
      "SAR": "97.7212"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
