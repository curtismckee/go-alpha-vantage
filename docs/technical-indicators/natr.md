<center>
  <h1>Technical Indicator - Normalized True Range</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the normalized average true range (NATR) values.


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=NATR` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time_period     | string  | true      | Number of data points used to calculate each NATR value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=NATR&symbol=MSFT&interval=weekly&time_period=14&apikey=demo](https://www.alphavantage.co/query?function=NATR&symbol=MSFT&interval=weekly&time_period=14&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Normalized Average True Range (NATR)",
    "3: Last Refreshed": "2019-03-14 12:50:40",
    "4: Interval": "weekly",
    "5: Time Period": 14,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: NATR": {
    "2019-03-14 12:50:40": {
      "NATR": "4.5471"
    },
    "2019-03-08": {
      "NATR": "4.7681"
    },
    "2019-03-01": {
      "NATR": "4.7385"
    },
    "2019-02-22": {
      "NATR": "5.0111"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
