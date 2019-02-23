<center>
  <h1>Search</h1>
</center>

##### Quick API Reference:

<!-- tabs:start -->

#### **Endpoint**

The Search Endpoint returns the best-matching symbols and market information based on keywords of your choice. The search results also contain match scores that provide you with the full flexibility to develop your own search and filtering logic. 

> GET `https://alphavantage.co/query?`

| Parameter       | Object  | Required  | Description |
| :---            | :---:   | :---:     | :---        |
| function        | string  | true      | The function of your choice. In this case, `function=SYMBOL_SEARCH` |
| keywords        | string  | true      | A text string of your choice. For example: `keywords=microsoft` |
| datatype        | string  | optional  | By default, `datatype=json`. Strings `json` and `csv` are accepted with the following specifications: `json` returns the search results in JSON format; `csv` returns the search results as a CSV (comma separated value) file. |
| apikey          | string  | true      | Your API key | 

Example (Click for JSON output):  
[https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=Micro&apikey=demo](https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=Micro&apikey=demo)

#### **Response**

> http code: 200

```javascript
{
  "bestMatches": [
    {
      "1. symbol": "AMD",
      "2. name": "Advanced Micro Devices Inc.",
      "3. type": "Equity",
      "4. region": "United States",
      "5. marketOpen": "09:30",
      "6. marketClose": "16:00",
      "7. timezone": "UTC-05",
      "8. currency": "USD",
      "9. matchScore": "0.5000"
    },
    {
      "1. symbol": "MU",
      "2. name": "Micron Technology Inc.",
      "3. type": "Equity",
      "4. region": "United States",
      "5. marketOpen": "09:30",
      "6. marketClose": "16:00",
      "7. timezone": "UTC-05",
      "8. currency": "USD",
      "9. matchScore": "0.4545"
    },
    {
      "1. symbol": "MSFT",
      "2. name": "Microsoft Corporation",
      "3. type": "Equity",
      "4. region": "United States",
      "5. marketOpen": "09:30",
      "6. marketClose": "16:00",
      "7. timezone": "UTC-05",
      "8. currency": "USD",
      "9. matchScore": "0.4444"
    },
    { ... },
    { ... },
    { ... }
  ]
}
```

> http code: 404

```javascript
{
    "Error Message": "Invalid API call. Please retry or visit the documentation (https://www.alphavantage.co/documentation/) for TIME_SERIES_INTRADAY."
}
```

<!-- tabs:end -->
