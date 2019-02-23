<center>
  <h1>Quotes</h1>
</center>

##### Quick API Reference:

<!-- tabs:start -->

### **Endpoint**

A lightweight alternative to the time series APIs, this service returns the latest price and volume information for a security of your choice.

> GET `https://alphavantage.co/query?function=GLOBAL_QUOTE`

| Parameter         | Object  | Required  | Description |
| :---              | :---:   | :---:     | :---        |
| function          | string  | true      | The function of your choice. In this case, `function=GLOBAL_QUOTE` |
| symbol            | string  | true      | The symbol of the global security of your choice. For example: `symbol=MSFT`
| datatype          | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the quote data in JSON format; `csv` returns the quote data as a CSV (comma separated value) file. |
| apikey          | string  | true      | Your API key | 

Example (Click for JSON output):  
[https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo](https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo)

### **Response**

> http code: 200

```javascript
{
  "Global Quote": {
    "01. symbol": "MSFT",
    "02. open": "107.8600",
    "03. high": "107.9400",
    "04. low": "106.2950",
    "05. price": "106.7900",
    "06. volume": "13085190",
    "07. latest trading day": "2019-02-20",
    "08. previous close": "107.7100",
    "09. change": "-0.9200",
    "10. change percent": "-0.8541%"
  }
}
```

> http code: 404

```javascript
{
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for TIME_SERIES_INTRADAY."
}
```

<!-- tabs:end -->
