<center>
  <h1>Technical Indicator - Money Flow Index</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the money flow index (MFI) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/03/072303.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=MoneyFlowIndex.htm)


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MFI` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each MFI value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=MFI&symbol=MSFT&interval=weekly&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=MFI&symbol=MSFT&interval=weekly&time_period=10&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Money Flow Index (MFI)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: MFI": {
    "2019-03-13": {
      "MFI": "75.7631"
    },
    "2019-03-08": {
      "MFI": "76.7142"
    },
    "2019-03-01": {
      "MFI": "73.7717"
    },
    "2019-02-22": {
      "MFI": "56.2683"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
