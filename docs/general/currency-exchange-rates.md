<center>
  <h1>Currency Exchange Rates</h1>
</center>

##### Quick API Reference:

<!-- tabs:start -->

### ** Endpoint **

This API returns the realtime exchange rate for any pair of digital currency (e.g. Bitcoin) or physical currency (e.g. USD). 

> GET `https://alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE`

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=CURRENCY_EXCHANGE_RATE` |
| from\_currency  | string  | true      | The currency you would like to get the exchange rate for. It can either be a physical currency or digital/crypto currency. For example: `from_currency=USD` or `from_currency=BTC` |  
| to\_currency    | string  | true      | The destination currency for the exchange rate. It can either be a physical currency or digital/crypto currency. For example: `to_currency=USD` or `to_currency=BTC` |
| apikey          | string  | true      | Your API key | 

Example (Click for JSON output):  
[https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo](https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo)


### **Response**

> http code: `200`  

```javascript
{
  "Realtime Currency Exchange Rate": {
    "1. From_Currency Code": "BTC",
    "2. From_Currency Name": "Bitcoin",
    "3. To_Currency Code": "CNY",
    "4. To_Currency Name": "Chinese Yuan",
    "5. Exchange Rate": "27513.09000000",
    "6. Last Refreshed": "2019-02-20 20:07:15",
    "7. Time Zone": "UTC"
  }
}
```

> http code: `404`

```javascript
{
  "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for CURRENCY_EXCHANGE_RATE."
}
```

<!-- tabs:end -->
