<center>
  <h1>Technical Indicator - Balance of Power</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the balance of power (BOP) values. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=BOP` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=BOP&symbol=MSFT&interval=daily&apikey=demo](https://www.alphavantage.co/query?function=BOP&symbol=MSFT&interval=daily&apikey=demo)


Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Balance Of Power (BOP)",
    "3: Last Refreshed": "2019-03-13 16:00:01",
    "4: Interval": "daily",
    "5: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: BOP": {
    "2019-03-13 16:00:01": {
      "BOP": "0.3033"
    },
    "2019-03-12": {
      "BOP": "0.5970"
    },
    "2019-03-11": {
      "BOP": "0.9340"
    },
    "2019-03-08": {
      "BOP": "0.7068"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
