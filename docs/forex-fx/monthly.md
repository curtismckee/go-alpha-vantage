<center>
  <h1>Forex -  Monthly</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the monthly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.  
The latest data point is the cumulative prices information for the month (or partial month) containing the current trading day, updated realtime. 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=FX_MONTHLY` |
| to\_symbol      | string  | true      | A three-letter symbol from the forex currency list. For example: `from_symbol=EUR`. |
| from\_symbol    | string  | true      | A three-letter symbol from the forex currency list. For example: `from_symbol=USD`. |
| datatype        | string  | true      | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the monthly time series in JSON format; `csv` returns the time series as a CSV (comma separated value) file. |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=FX_MONTHLY&from_symbol=EUR&to_symbol=USD&apikey=demo](https://www.alphavantage.co/query?function=FX_MONTHLY&from_symbol=EUR&to_symbol=USD&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "1. Information": "Forex Monthly Prices (open, high, low, close)",
    "2. From Symbol": "EUR",
    "3. To Symbol": "USD",
    "4. Last Refreshed": "2019-03-08 05:20:00",
    "5. Time Zone": "GMT+8"
  },
  "Time Series FX (Monthly)": {
    "2019-03-08": {
      "1. open": "1.1371",
      "2. high": "1.1409",
      "3. low": "1.1175",
      "4. close": "1.1186"
    },
    "2019-02-28": {
      "1. open": "1.1446",
      "2. high": "1.1488",
      "3. low": "1.1233",
      "4. close": "1.1371"
    },
    "2019-01-31": {
      "1. open": "1.1465",
      "2. high": "1.1570",
      "3. low": "1.1288",
      "4. close": "1.1446"
    },
    { ... },
    { ... },
    { ... }
  }
}
```

<!-- tabs:end -->
