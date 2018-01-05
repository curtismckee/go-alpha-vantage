package main

import (
	"flag"
	"github.com/cmckee-dev/go-alpha-vantage"
	"fmt"
	"os"
)

var (
	apiKey = flag.String("apikey", "", "api key for alpha vantage")
	symbol = flag.String("symbol", "GOOGL", "symbol to list")
)

func main() {
	flag.Parse()

	client := av.NewClient(*apiKey)

	runStuff(client, *symbol, av.TimeSeriesDaily)
	runStuff(client, *symbol, av.TimeSeriesWeekly)
	runStuff(client, *symbol, av.TimeSeriesMonthly)
}

func runStuff(client *av.Client, symbol string, series av.TimeSeries) {
	res, err := client.StockTimeSeries(series, symbol)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading series %s: %v", series, err)
	}

	fmt.Printf("%s %s\n", series, symbol)
	for _, record := range res {
		fmt.Printf("%s: high=%f low=%f\n", record.Date, record.High, record.Low)
	}
}
