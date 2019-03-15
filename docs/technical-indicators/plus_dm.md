<center>
  <h1>Technical Indicator - Plus Directional Movement</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the plus directional movement (PLUS_DM) values. See also: [Investopedia article](https://www.investopedia.com/articles/technical/02/050602.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=PLUS_DM` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each PLUS_DM value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=PLUS_DM&symbol=MSFT&interval=daily&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=PLUS_DM&symbol=MSFT&interval=daily&time_period=10&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Plus Directional Movement (PLUS_DM)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: PLUS_DM": {
    "2019-03-13": {
      "PLUS_DM": "6.5638"
    },
    "2019-03-12": {
      "PLUS_DM": "6.1709"
    },
    "2019-03-11": {
      "PLUS_DM": "5.7010"
    },
    "2019-03-08": {
      "PLUS_DM": "3.8455"
    },
    "2019-03-07": {
      "PLUS_DM": "4.2728"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
