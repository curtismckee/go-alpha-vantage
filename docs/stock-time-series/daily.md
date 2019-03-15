<center>
  <h1>Time Series - Daily</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume) of the global equity specified, covering 20+ years of historical data.  
The most recent data point is the cumulative prices and volume information of the current trading day, updated realtime.  

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=TIME_SERIES_INTRADAY` |
| symbol          | string  | true      | The name of the equity of your choice. The example: `symbol=MSFT` |
| interval        | string  | true      | Time interval between two consecutive data points in the time series. The following calues are supported: `1min`, `5min`, `15min`, `30min`, `60min` |
| outputsize      | string  | optional  | By default, `outputsize=compact`. Strings `compact` and `full` are accepted with the following specifications: `compact` returns only the latest 100 data points; `full` returns the full-length intraday time series. The "compact" option is recommended if you would like to reduce the data size of each API call.
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the intraday times series in JSON format; `csv` returns the time seris as a CSV (comma spearated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=MSFT&apikey=demo](https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=MSFT&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1. Information": "Daily Prices (open, high, low, close) and Volumes",
    "2. Symbol": "MSFT",
    "3. Last Refreshed": "2019-02-22",
    "4. Output Size": "Compact",
    "5. Time Zone": "US/Eastern"
  },
  "Time Series (Daily)": {
    "2019-02-22": {
    "1. open": "110.0500",
    "2. high": "111.2000",
    "3. low": "109.8200",
    "4. close": "110.9700",
    "5. volume": "27763218"
  },
  { ... },
  { ... },
  { ... }
}
```

<!-- tabs:end -->
