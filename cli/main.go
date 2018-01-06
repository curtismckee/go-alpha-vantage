package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cmckee-dev/go-alpha-vantage"
	"github.com/kr/pretty"
)

const (
	physicalCurrency = "USD"
)

var (
	apiKey       = flag.String("apikey", "", "api key for alpha vantage")
	symbol       = flag.String("symbol", "GOOGL", "symbol to list")
	cryptoSymbol = flag.String("crypto", "ETH", "crypto-currency to query")
)

func main() {
	flag.Parse()

	client := av.NewClient(*apiKey)

	queryCrypto(client, *cryptoSymbol, physicalCurrency)

	for i := av.TimeIntervalOneMinute; i <= av.TimeIntervalSixtyMinute; i++ {
		queryInterval(client, *symbol, i)
	}

	for t := av.TimeSeriesDaily; t <= av.TimeSeriesMonthlyAdjusted; t++ {
		queryTimeSeries(client, *symbol, t)
	}
}

func queryTimeSeries(client *av.Client, symbol string, series av.TimeSeries) {
	res, err := client.StockTimeSeries(series, symbol)
	if err != nil {
		ErrorF("error loading series %s: %v", series, err)
		return
	}
	fmt.Printf("%s %s with %d records\n", series, symbol, len(res))
}

func queryInterval(client *av.Client, symbol string, timeInterval av.TimeInterval) {
	res, err := client.StockTimeSeriesIntraday(timeInterval, symbol)
	if err != nil {
		ErrorF("error loading intraday series %s: %v", timeInterval, err)
		return
	}
	fmt.Printf("%s %s with %d records\n", timeInterval, symbol, len(res))
}

func queryCrypto(client *av.Client, digital, physical string) {
	res, err := client.DigitalCurrency(digital, physical)
	if err != nil {
		ErrorF("error loading crypto: %s => %s: %v", digital, physical, err)
		return
	}
	for _, record := range res {
		pretty.Println(record)
	}
	fmt.Printf("%s => %s with %d records\n", digital, physical, len(res))
}

func ErrorF(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", format), args...)
}
