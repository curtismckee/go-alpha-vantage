<center>
  <h1>Time Series - Weekly</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns weekly time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly volume) of the global equity specified, covering 20+ years of historical data.  
The latest data point is the cumulative prices and volume information for the week (or partial week) that contains the current trading day, updated realtime. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=TIME_SERIES_INTRADAY` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the weekly times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=MSFT&apikey=demo](https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=MSFT&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1. Information": "Weekly Prices (open, high, low, close) and Volumes",
    "2. Symbol": "MSFT",
    "3. Last Refreshed": "2019-02-25",
    "4. Time Zone": "US/Eastern"
  },
  "Weekly Time Series": {
    "2019-02-25": {
      "1. open": "111.7600",
      "2. high": "112.1800",
      "3. low": "111.2600",
      "4. close": "111.5900",
      "5. volume": "23750599"
    },
    { ... },
    { ... },
    { ... },
  }
}
```

<!-- tabs:end -->
