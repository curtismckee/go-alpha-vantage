<center>
  <h1>Technical Indicator - Ultimate Oscillator</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**


This API returns the ultimate oscillator (ULTOSC) values. See also: [mathematical reference](https://www.fmlabs.com/reference/default.htm?url=UltimateOsc.htm)


| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=ULTOSC` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| timeperiod1     | string  | optional  | The first time period for the indicator. Positive integers are accepted. By default, `timeperiod1=7`. |
| timeperiod2     | string  | optional  | The second time period for the indicator. Positive integers are accepted. By default, `timeperiod2=14`. |
| timeperiod3     | string  | optional  | The third time period for the indicator. Positive integers are accepted. By default, `timeperiod3=28`. |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=ULTOSC&symbol=MSFT&interval=daily&timeperiod1=8&apikey=demo](https://www.alphavantage.co/query?function=ULTOSC&symbol=MSFT&interval=daily&timeperiod1=8&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Ultimate Oscillator (ULTOSC)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "daily",
    "5.1: Time Period 1": 8,
    "5.2: Time Period 2": 14,
    "5.3: Time Period 3": 28,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: ULTOSC": {
    "2019-03-13": {
      "ULTOSC": "63.1844"
    },
    "2019-03-12": {
      "ULTOSC": "63.9983"
    },
    "2019-03-11": {
      "ULTOSC": "61.4522"
    },
    "2019-03-08": {
      "ULTOSC": "57.6948"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
