<center>
  <h1>Technical Indicator - Average Directional Movement Index Rating</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the average directional movement index rating (ADXR) values. See also: [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=ADXR.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=ADXR` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each ADXR value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=ADXR&symbol=MSFT&interval=daily&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=ADXR&symbol=MSFT&interval=daily&time_period=10&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Average Directional Movement Index Rating (ADXR)",
    "3: Last Refreshed": "2019-03-13 16:00:01",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: ADXR": {
    "2019-03-13 16:00:01": {
      "ADXR": "21.1802"
    },
    "2019-03-12": {
      "ADXR": "19.5403"
    },
    "2019-03-11": {
      "ADXR": "18.3374"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
