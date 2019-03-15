<center>
  <h1>Technical Indicator - True Range</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

[This API returns the true range (TRANGE) values. See also: mathematical reference](This API returns the true range (TRANGE) values. See also: mathematical reference)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=TRANGE` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=TRANGE&symbol=MSFT&interval=daily&apikey=demo](https://www.alphavantage.co/query?function=TRANGE&symbol=MSFT&interval=daily&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "True Range (TRANGE)",
    "3: Last Refreshed": "2019-03-14 11:58:50",
    "4: Interval": "daily",
    "5: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: TRANGE": {
    "2019-03-14 11:58:50": {
      "TRANGE": "0.7500"
    },
    "2019-03-13": {
      "TRANGE": "1.3800"
    },
    "2019-03-12": {
      "TRANGE": "1.3401"
    },
    "2019-03-11": {
      "TRANGE": "2.4400"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
