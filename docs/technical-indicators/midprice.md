<center>
  <h1>Technical Indicator - Midprice Values</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the midpoint price (MIDPRICE) values. MIDPRICE = (highest high + lowest low)/2.

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=MIDPRICE` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each MIDPRICE value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=MIDPRICE&symbol=MSFT&interval=daily&time_period=10&apikey=demo](https://www.alphavantage.co/query?function=MIDPRICE&symbol=MSFT&interval=daily&time_period=10&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Midpoint Price over period (MIDPRICE)",
    "3: Last Refreshed": "2019-03-13",
    "4: Interval": "daily",
    "5: Time Period": 10,
    "6: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: MIDPRICE": {
    "2019-03-13": {
      "MIDPRICE": "111.9000"
    },
    "2019-03-12": {
      "MIDPRICE": "111.3950"
    },
    "2019-03-11": {
      "MIDPRICE": "111.0250"
    },
    "2019-03-08": {
      "MIDPRICE": "111.0250"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
