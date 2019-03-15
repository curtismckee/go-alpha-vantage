<center>
  <h1>Technical Indicator - Minus Directional Indicator</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

s API returns the minus directional indicator (MINUS_DI) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/02/050602.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=DI.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MINUS_DI` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each MINUS_DI value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=MINUS_DI&symbol=MSFT&interval=weekly&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=MINUS_DI&symbol=MSFT&interval=weekly&time_period=10&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Minus Directional Indicator (MINUS_DI)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: MINUS_DI": {
    "2019-03-13": {
      "MINUS_DI": "18.5386"
    },
    "2019-03-08": {
      "MINUS_DI": "20.3267"
    },
    "2019-03-01": {
      "MINUS_DI": "17.8402"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
