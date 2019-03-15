<center>
  <h1>Technical Indicators - Moving Average Convergence / Divergence with Controllable Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**


This API returns the moving average convergence / divergence values with controllable moving average type. See also: [Investopedia article](https://www.investopedia.com/articles/forex/05/macddiverge.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=MACD.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MACDEXT` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| fastperiod      | string  | optional  | Positive integers are accepted. By default, `fastperiod=12`. |
| slowperiod      | string  | optional  | Positive integers are accepted. By default, `slowperiod=26`. |
| signalperiod    | string  | optional  | Positive integers are accepted. By default, `signalperiod=9`. |
| fastmatype      | string  | optional  | Moving average type for the faster moving average. By default, `fastmatype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| slowmatype      | string  | optional  | Moving average type for the slower moving average. By default, `slowmatype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| signalmatype    | string  | optional  | Moving average type for the signal moving average. By default, `signalmatype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=MACDEXT&symbol=MSFT&interval=daily&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=MACDEXT&symbol=MSFT&interval=daily&series_type=open&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "MACD with Controllable MA Type (MACDEXT)",
    "3: Last Refreshed": "2019-03-12",
    "4: Interval": "daily",
    "5.1: Fast Period": 12,
    "5.2: Slow Period": 26,
    "5.3: Signal Period": 9,
    "5.4: Fast MA Type": 0,
    "5.5: Slow MA Type": 0,
    "5.6: Signal MA Type": 0,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: MACDEXT": {
    "2019-03-12": {
      "MACD_Signal": "2.6246",
      "MACD_Hist": "0.1643",
      "MACD": "2.7888"
    },
    "2019-03-11": {
      "MACD_Signal": "2.4961",
      "MACD_Hist": "0.4098",
      "MACD": "2.9059"
    },
    "2019-03-08": {
      "MACD_Signal": "2.3052",
      "MACD_Hist": "0.5364",
      "MACD": "2.8416"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
