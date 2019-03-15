<center>
  <h1>Technical Indicator - Aroon</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**


This API returns the Aroon (AROON) values. See also: [Investopedia article](https://www.investopedia.com/articles/trading/06/aroon.asp) and [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=Aroon.htm)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=AROON` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each AROON value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=AROON&symbol=MSFT&interval=daily&time_period=14&apikey=demo](https://www.alphavantage.co/query?function=AROON&symbol=MSFT&interval=daily&time_period=14&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Aroon (AROON)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "daily",
    "5: Time Period": 14,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: AROON": {
    "2019-03-13": {
      "Aroon Down": "0.0000",
      "Aroon Up": "100.0000"
    },
    "2019-03-12": {
      "Aroon Down": "0.0000",
      "Aroon Up": "100.0000"
    },
    "2019-03-11": {
      "Aroon Down": "7.1429",
      "Aroon Up": "64.2857"
    },
    "2019-03-08": {
      "Aroon Down": "14.2857",
      "Aroon Up": "71.4286"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
