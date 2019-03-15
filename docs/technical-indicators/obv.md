<center>
  <h1>Technical Indicator - On Balance Volume</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

s API returns the on balance volume (OBV) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/100801.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=OBV.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=OBV` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=OBV&symbol=MSFT&interval=weekly&apikey=demo](https://www.alphavantage.co/query?function=OBV&symbol=MSFT&interval=weekly&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "On Balance Volume (OBV)",
    "3: Last Refreshed": "2019-03-14 13:30:43",
    "4: Interval": "weekly",
    "5: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: OBV": {
    "2019-03-14 13:30:43": {
      "OBV": "211365872.0000"
    },
    "2019-03-08": {
      "OBV": "112667632.0000"
    },
    "2019-03-01": {
      "OBV": "224658344.0000"
    },
    "2019-02-22": {
      "OBV": "105298847.0000"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
