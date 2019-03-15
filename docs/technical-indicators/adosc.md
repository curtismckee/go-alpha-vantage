<center>
  <h1>Technical Indicator - Chaikin A/D Oscillator</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Chaikin A/D oscillator (ADOSC) values. See also: [Investopedia article](https://www.investopedia.com/articles/active-trading/031914/understanding-chaikin-oscillator.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=AccumDist.htm)


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=ADOSC` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| fastperiod      | string  | optional  | The time period of the fast EMA. Positive integers are accepted. By default, `fastperiod=3`. |
| slowperiod      | string  | optional  | The time period of the slow EMA. Positive integers are accepted. By default, `slowperiod=10`. |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=ADOSC&symbol=MSFT&interval=daily&fastperiod=5&apikey=demo](https://www.alphavantage.co/query?function=ADOSC&symbol=MSFT&interval=daily&fastperiod=5&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Chaikin A/D Oscillator (ADOSC)",
    "3: Last Refreshed": "2019-03-14 13:13:23",
    "4: Interval": "daily",
    "5.1: FastK Period": 5,
    "5.2: SlowK Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: ADOSC": {
    "2019-03-14 13:13:23": {
      "ADOSC": "15125458.4369"
    },
    "2019-03-13": {
      "ADOSC": "14296223.4649"
    },
    "2019-03-12": {
      "ADOSC": "12436498.6044"
    },
    "2019-03-11": {
      "ADOSC": "9413931.7793"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
