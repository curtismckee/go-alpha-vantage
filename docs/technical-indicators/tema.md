<center>
  <h1>Technical Indicators - Triple Exponential Moving Average</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**


This API returns the triple exponential moving average (TEMA) values. See also: [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=TEMA.htm)


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicators of your choice. In this case, `function=TEMA` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each moving average value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=TEMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo](https://www.alphavantage.co/query?function=TEMA&symbol=MSFT&interval=weekly&time_period=10&series_type=open&apikey=demo) 


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Triple Exponential Moving Average (TEMA)",
    "3: Last Refreshed": "2019-03-11",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Series Type": "open",
    "7: Time Zone": "US/Eastern"
  },
  "Technical Analysis: TEMA": {
    "2019-03-11": {
      "TEMA": "111.8031"
    },
    "2019-03-08": {
      "TEMA": "111.1993"
    },
    "2019-03-01": {
      "TEMA": "108.7999"
    },
    "2019-02-22": {
      "TEMA": "106.0005"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
