<center>
  <h1>Sector Preformances</h1>
</center>

<!-- tabs:start -->

### **Client**

Coming Soon!

### **API Reference**

This API returns the realtime and historical sector performances calculated from S&P500 incumbents.

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=SECTOR` |
| apikey          | string  | true      | Your API key | 

Example JSON Endpoint:  

[https://www.alphavantage.co/query?function=SECTOR&apikey=demo](https://www.alphavantage.co/query?function=SECTOR&apikey=demo)

Example Response:  

```javascript
{
  "Meta Data": {
    "Information": "US Sector Performance (realtime & historical)",
    "Last Refreshed": "04:20 PM ET 02/22/2019"
  },
  "Rank A: Real-Time Performance": {
    "Information Technology": "1.29%",
    "Communication Services": "1.05%",
    "Health Care": "0.93%",
    "Industrials": "0.64%",
    "Utilities": "0.61%",
    "Real Estate": "0.59%",
    "Consumer Discretionary": "0.58%",
    "Materials": "0.41%",
    "Energy": "0.33%",
    "Financials": "-0.23%",
    "Consumer Staples": "-0.28%"
  },
  "Rank B: 1 Day Performance": {
    "Information Technology": "1.29%",
    "Communication Services": "1.05%",
    "Health Care": "0.93%",
    "Industrials": "0.64%",
    "Utilities": "0.61%",
    "Real Estate": "0.59%",
    "Consumer Discretionary": "0.58%",
    "Materials": "0.41%",
    "Energy": "0.33%",
    "Financials": "-0.23%",
    "Consumer Staples": "-0.28%"
  },
  { ... }.
  { ... },
  { ... }
}
```

<!-- tabs:end -->
