<center>
  <h1>Technical Indicators - Volume Weighted Average Price</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the volume weighted average price (VWAP) for intraday time series. See also: [Investopedia article](https://www.investopedia.com/terms/v/vwap.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=VWAP` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `weekly`, `monthly` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=VWAP&symbol=MSFT&interval=15min&apikey=demo](https://www.alphavantage.co/query?function=VWAP&symbol=MSFT&interval=15min&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Volume Weighted Average Price (VWAP)",
    "3: Last Refreshed": "2019-03-12 12:45:00",
    "4: Interval": "15min",
    "5: Time Zone": "US/Eastern"
  },
  "Technical Analysis: VWAP": {
    "2019-03-12 12:45": {
      "VWAP": "113.4578"
    },
    "2019-03-12 12:30": {
      "VWAP": "113.4462"
    },
    "2019-03-12 12:15": {
      "VWAP": "113.4244"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
