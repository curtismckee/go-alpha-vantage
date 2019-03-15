<center>
  <h1>Technical Indicator - Hillbert Transform, Sine Wave</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Hilbert transform, sine wave (HT_SINE) values.

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=HT_SINE` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series_type     | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=HT_SINE&symbol=MSFT&interval=daily&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=HT_SINE&symbol=MSFT&interval=daily&series_type=close&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Hilbert Transform - SineWave (HT_SINE)",
    "3: Last Refreshed": "2019-03-14 13:49:09",
    "4: Interval": "daily",
    "5: Series Type": "close",
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: HT_SINE": {
    "2019-03-14 13:49:09": {
      "LEAD SINE": "-0.9304",
      "SINE": "-0.3986"
    },
    "2019-03-13": {
      "LEAD SINE": "-0.9107",
      "SINE": "-0.3519"
    },
    "2019-03-12": {
      "LEAD SINE": "-0.8818",
      "SINE": "-0.2900"
    },
    "2019-03-11": {
      "LEAD SINE": "-0.8610",
      "SINE": "-0.2492"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
