<center>
  <h1>Technical Indicator - Bollinger Bands</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the Bollinger bands (BBANDS) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/04/030304.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=Bollinger.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=BBANDS` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each BBANDS value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| nbdevup         | string  | optional  | The standard deviation multiplier of the upper band. Positive integers are accepted. By default, `nbdevup=2`. |
| nbdevdn         | string  | optional  | The standard deviation multiplier of the lower band. Positive integers are accepted. By default, `nbdevdn=2`. |
| matype          | string  | optional  | Moving average type of the time series. By default, `matype=0`. Integers 0 - 8 are accepted with the following mappings. 0 = Simple Moving Average (SMA), 1 = Exponential Moving Average (EMA), 2 = Weighted Moving Average (WMA), 3 = Double Exponential Moving Average (DEMA), 4 = Triple Exponential Moving Average (TEMA), 5 = Triangular Moving Average (TRIMA), 6 = T3 Moving Average, 7 = Kaufman Adaptive Moving Average (KAMA), 8 = MESA Adaptive Moving Average (MAMA). |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=BBANDS&symbol=MSFT&interval=weekly&time_period=5&series_type=close&nbdevup=3&nbdevdn=3&apikey=demo](https://www.alphavantage.co/query?function=BBANDS&symbol=MSFT&interval=weekly&time_period=5&series_type=close&nbdevup=3&nbdevdn=3&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Bollinger Bands (BBANDS)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "weekly",
    "5: Time Period": 5,
    "6.1: Deviation multiplier for upper band": 3,
    "6.2: Deviation multiplier for lower band": 3,
    "6.3: MA Type": 0,
    "7: Series Type": "close",
    "8: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: BBANDS": {
    "2019-03-13": {
      "Real Upper Band": "117.6334",
      "Real Lower Band": "105.0586",
      "Real Middle Band": "111.3460"
    },
    "2019-03-08": {
      "Real Upper Band": "116.7596",
      "Real Lower Band": "102.4004",
      "Real Middle Band": "109.5800"
    },
    "2019-03-01": {
      "Real Upper Band": "118.6033",
      "Real Lower Band": "97.4647",
      "Real Middle Band": "108.0340"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
