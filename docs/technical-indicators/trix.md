<center>
  <h1>Technical Indicator - 1 Day Rate of Change of Triple Smooth Exponential Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the 1-day rate of change of a triple smooth exponential moving average (TRIX) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/02/092402.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=TRIX.htm)


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=TRIX` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each TRIX value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=TRIX&symbol=MSFT&interval=daily&time_period=10&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=TRIX&symbol=MSFT&interval=daily&time_period=10&series_type=close&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "1-day Rate-Of-Change (ROC) of a Triple Smooth EMA (TRIX)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6: Series Type": "close",
    "7: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: TRIX": {
    "2019-03-13": {
      "TRIX": "0.2517"
    },
    "2019-03-12": {
      "TRIX": "0.2493"
    },
    "2019-03-11": {
      "TRIX": "0.2561"
    },
    "2019-03-08": {
      "TRIX": "0.2725"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
