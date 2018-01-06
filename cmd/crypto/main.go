// a program to check the value of a crypto currency
package main

import (
	"flag"
	"fmt"
	"strings"
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

	for _, symbol := range strings.Split(*cryptoSymbols, ",") {
		queryCrypto(client, symbol, *physicalCurrency)
	}
}

func queryCrypto(client *av.Client, digital, physical string) {
	fmt.Printf("%s = ", digital)
	res, err := client.DigitalCurrency(digital, physical)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	if len(res) == 0 {
		fmt.Print("no data\n")
		return
	}
	record := res[len(res)-1]
	fmt.Printf("%f %s (updated %s)\n", record.Price, physical, asLocalTime(record.Time))
}

func asLocalTime(in time.Time) string {
	return in.In(time.Local).Format(timeFormat)
}
