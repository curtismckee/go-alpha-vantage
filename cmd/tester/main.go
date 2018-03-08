package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/cmckee-dev/go-alpha-vantage"
)

var (
	apiKey           = flag.String("apikey", "", "api key for alpha vantage")
	symbol           = flag.String("symbol", "GOOGL", "symbol to list")
	cryptoSymbol     = flag.String("crypto", "ETH", "crypto-currency to query")
	physicalCurrency = flag.String("dollars", "USD", "physical currency to query value of crypto")
)

func main() {
	flag.Parse()

	client := av.NewClient(*apiKey)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		queryCrypto(client, *cryptoSymbol, *physicalCurrency)
	}()

	for interval := av.TimeIntervalOneMinute; interval <= av.TimeIntervalSixtyMinute; interval++ {
		wg.Add(1)
		go func(interval av.TimeInterval) {
			defer wg.Done()
			queryInterval(client, *symbol, interval)
		}(interval)
	}

	for series := av.TimeSeriesDaily; series <= av.TimeSeriesMonthlyAdjusted; series++ {
		wg.Add(1)
		go func(series av.TimeSeries) {
			defer wg.Done()
			queryTimeSeries(client, *symbol, series)
		}(series)
	}

	wg.Wait()
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
	fmt.Printf("%s => %s with %d records\n", digital, physical, len(res))
}

func ErrorF(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", format), args...)
}
