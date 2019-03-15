<center>
  <h1>Exchange Rates</h1>
</center>

<!-- tabs:start -->
### **Client**

Coming Soon!

### **API Reference**

This API returns the realtime exchange rate for any pair of digital currency (e.g. Bitcoin) or physical currency (e.g. USD). 

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=CURRENCY_EXCHANGE_RATE` |
| from\_currency  | string  | true      | The currency you would like to get the exchange rate for. It can either be a physical currency or digital/crypto currency. For example: `from_currency=USD` or `from_currency=BTC` |  
| to\_currency    | string  | true      | The destination currency for the exchange rate. It can either be a physical currency or digital/crypto currency. For example: `to_currency=USD` or `to_currency=BTC` |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo](https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo)

Example Response:

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
<!-- tabs:end -->
