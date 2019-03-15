<center>
  <h1>Technical Indicator - Hillbert Transform, Phasor Components</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Hilbert transform, phasor components (HT_PHASOR) values. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=HT_PHASOR` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series_type     | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=HT_PHASOR&symbol=MSFT&interval=weekly&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=HT_PHASOR&symbol=MSFT&interval=weekly&series_type=close&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Hilbert Transform - Phasor Components (HT_PHASOR)",
    "3: Last Refreshed": "2019-03-14 16:00:02",
    "4: Interval": "weekly",
    "5: Series Type": "close",
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: HT_PHASOR": {
    "2019-03-14 16:00:02": {
      "QUADRATURE": "2.3242",
      "PHASE": "0.7289"
    },
    "2019-03-08": {
      "QUADRATURE": "-9.1717",
      "PHASE": "2.4184"
    },
    "2019-03-01": {
      "QUADRATURE": "-2.4937",
      "PHASE": "7.5829"
    },
    "2019-02-22": {
      "QUADRATURE": "13.4479",
      "PHASE": "5.2225"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
