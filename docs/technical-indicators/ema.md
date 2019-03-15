<center>
  <h1>Technical Indicators - Exponential Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the exponential moving average (EMA) values. See also: [mathematical reference].

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=EMA` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_series    | string  | true      | Number of data points used to calculate each moving average value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`)
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=EMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=EMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Exponential Moving Average (EMA)",
    "3: Last Refreshed": "2019-03-11 16:00:01",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern"
  },
  "Technical Analysis: EMA": {
    "2019-03-11 16:00:01": {
      "EMA": "108.2846"
    },
    "2019-03-08": {
      "EMA": "107.6834"
    },
    "2019-03-01": {
      "EMA": "106.4975"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
