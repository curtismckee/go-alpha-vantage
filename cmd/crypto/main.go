// a program to check the value of a crypto currency
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cmckee-dev/go-alpha-vantage"
)

var (
	apiKey           = flag.String("apikey", "", "api key for alpha vantage")
	cryptoSymbol     = flag.String("crypto", "ETH", "crypto-currency to query")
	physicalCurrency = flag.String("dollar", "USD", "physical currency to query value of crypto")
)

func main() {
	flag.Parse()
	client := av.NewClient(*apiKey)
	fmt.Printf("querying %s price in %s...\n", *cryptoSymbol, *physicalCurrency)
	queryCrypto(client, *cryptoSymbol, *physicalCurrency)
}

func queryCrypto(client *av.Client, digital, physical string) {
	res, err := client.DigitalCurrency(digital, physical)
	if err != nil {
		ExitF(2, "could not query %s => %s: %v", digital, physical, err)
	}
	if len(res) == 0 {
		ExitF(3, "no records found for %s => %s", digital, physical)
	}
	record := res[len(res)-1]
	fmt.Printf("%s: 1 %s = %f %s\n", record.Time, digital, record.Price, physical)
}

func ErrorF(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", format), args...)
}
func ExitF(code int, format string, args ...interface{}) {
	ErrorF(format, args...)
	os.Exit(code)
}
