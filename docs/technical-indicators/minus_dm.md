<center>
  <h1>Technical Indicator - Minus Directional Movement</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the minus directional movement (MINUS\_DM) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/02/050602.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MINUS_DM` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each MINUS_DM value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=MINUS_DM&symbol=MSFT&interval=daily&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=MINUS_DM&symbol=MSFT&interval=daily&time_period=10&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Minus Directional Movement (MINUS_DM)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: MINUS_DM": {
    "2019-03-13": {
      "MINUS_DM": "3.3415"
    },
    "2019-03-12": {
      "MINUS_DM": "3.7128"
    },
    "2019-03-11": {
      "MINUS_DM": "4.1253"
    },
    "2019-03-08": {
      "MINUS_DM": "4.5837"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
