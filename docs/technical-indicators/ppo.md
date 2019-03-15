<center>
  <h1>Technical Indicator - Percentage Price Oscillator</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the percentage price oscillator (PPO) values. See also: [Investopedia article](https://www.investopedia.com/articles/investing/051214/use-percentage-price-oscillator-elegant-indicator-picking-stocks.asp) and [mathematical reference](https://www.investopedia.com/articles/investing/051214/use-percentage-price-oscillator-elegant-indicator-picking-stocks.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=PPO` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| fastperiod      | string  | optional  | Positive integers are accepted. By default, `fastperiod=12`. |
| slowperiod      | string  | optional  | Positive integers are accepted. By default, `slowperiod=26`. |
| matype          | string  | optional  | Moving average type. By default, `matype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 


Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=PPO&symbol=MSFT&interval=daily&series_type=close&fastperiod=10&matype=1&apikey=demo](https://www.alphavantage.co/query?function=PPO&symbol=MSFT&interval=daily&series_type=close&fastperiod=10&matype=1&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Percentage Price Oscillator (PPO)",
    "3: Last Refreshed": "2019-03-13 16:00:01",
    "4: Interval": "daily",
    "5.1: Fast Period": 10,
    "5.2: Slow Period": 26,
    "5.3: MA Type": 1,
    "6: Series Type": "close",
    "7: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: PPO": {
    "2019-03-13 16:00:01": {
      "PPO": "1.8728"
    },
    "2019-03-12": {
      "PPO": "1.7343"
    },
    "2019-03-11": {
      "PPO": "1.6313"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
