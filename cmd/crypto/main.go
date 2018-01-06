// a program to check the value of a crypto currency
package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cmckee-dev/go-alpha-vantage"
)

const (
	timeFormat = "2006/01/02 3:04pm"
)

var (
	apiKey           = flag.String("apikey", "", "api key for alpha vantage")
	cryptoSymbols    = flag.String("crypto", "BTC,ETH,LTC,TRX", "comma separated crypto-currencies to query")
	physicalCurrency = flag.String("dollar", "USD", "physical currency to query value of crypto")
)

func main() {
	flag.Parse()
	client := av.NewClient(*apiKey)
	fmt.Printf("querying price in %s...\n", *physicalCurrency)

	wg := &sync.WaitGroup{}

	for _, symbol := range strings.Split(*cryptoSymbols, ",") {
		wg.Add(1)
		go func(symbol string) {
			defer wg.Done()
			queryCrypto(client, symbol, *physicalCurrency)
		}(symbol)
	}

	wg.Wait()
}

func queryCrypto(client *av.Client, digital, physical string) {
	res, err := client.DigitalCurrency(digital, physical)
	if err != nil {
		fmt.Printf("%s error: %v\n", digital, err)
		return
	}
	if len(res) == 0 {
		fmt.Printf("%s = no data\n", digital)
		return
	}
	record := res[len(res)-1]
	fmt.Printf("%s = %f %s (updated %s)\n", digital, record.Price, physical, asLocalTime(record.Time))
}

func asLocalTime(in time.Time) string {
	return in.In(time.Local).Format(timeFormat)
}
