<center>
  <h1>Technical Indicator - Rate of Change</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the rate of change (ROC) values. See also: [Investopedia article](https://www.investopedia.com/terms/p/pricerateofchange.asp)

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The technical indicator of your choice. In this case, `function=ROC` |
| symbol          | string  | true      | The name of the security of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min`, `daily`, `weekly`, `monthly` |
| time\_period    | string  | true      | Number of data points used to calculate each ROC value. Positive integers are accepted (e.g., `time_period=60`, `time_period=200`) |
| series\_type    | string  | true      | The desired price type in the time series. Four types are supported: `close`, `open`, `high`, `low` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the daily time series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  


[https://www.alphavantage.co/query?function=ROC&symbol=MSFT&interval=weekly&time_period=10&series_type=close&apikey=demo](https://www.alphavantage.co/query?function=ROC&symbol=MSFT&interval=weekly&time_period=10&series_type=close&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1: Symbol": "MSFT",
    "2: Indicator": "Rate of change : ((price/prevPrice)-1)*100",
    "3: Last Refreshed": "2019-03-13 16:00:01",
    "4: Interval": "weekly",
    "5: Time Period": 10,
    "6: Series Type": "close",
    "7: Time Zone": "US/Eastern Time"
  },
  "Technical Analysis: ROC": {
    "2019-03-13 16:00:01": {
      "ROC": "12.3320"
    },
    "2019-03-08": {
      "ROC": "10.0807"
    },
    "2019-03-01": {
      "ROC": "14.5577"
    },
    "2019-02-22": {
      "ROC": "4.6591"
    },
    "2019-02-15": {
      "ROC": "3.2437"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
