<center>
  <h1>Technical Indicators - Stochastic Fast</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the stochastic fast (STOCHF) values. See also: [Investopedia article](https://www.investopedia.com/university/indicator_oscillator/ind_osc8.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=StochasticOscillator.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=STOCHF` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| fastkperiod     | string  | optional  | The time period of the fastk moving average. Positive integers are accepted. By default, `fastkperiod=5`. |
| fastdperiod     | string  | optional  | The time period of the fastd moving average. Positive integers are accepted. By default, `fastdperiod=3`. |
| fastdmatype     | string  | optional  | Moving average type for the fastd moving average. By default, `fastdmatype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=STOCHF&symbol=MSFT&interval=daily&apikey=demo](https://www.alphavantage.co/query?function=STOCHF&symbol=MSFT&interval=daily&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Stochastic Fast (STOCHF)",
    "3: Last Refreshed": "2019-03-12",
    "4: Interval": "daily",
    "5.1: FastK Period": 5,
    "5.2: FastD Period": 3,
    "5.3: FastD MA Type": 0,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: STOCHF": {
    "2019-03-12": {
      "FastK": "92.8709",
      "FastD": "76.1354"
    },
    "2019-03-11": {
      "FastK": "97.1084",
      "FastD": "50.3483"
    },
    "2019-03-08": {
      "FastK": "38.4270",
      "FastD": "30.9040"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
