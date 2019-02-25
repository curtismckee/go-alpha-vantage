package av

import (
	"encoding/csv"
	"github.com/pkg/errors"
	"io"
	"time"
)

// Quotes specifies a given security to query for.
type Quotes string

const (
	GlobalQuote = "GLOBAL_QUOTE"
)

var (
	// quotesDateFormats are the expected date formats in quotes data
	quotesDateFormats = []string{
		"2006-01-02",
	}
)

// TimeSeriesValue is a piece of data for a given time about stock prices
type QuoteValue struct {
	Symbol        string
	Open          float64
	High          float64
	Low           float64
	Price         float64
	Volume        int
	LatestDay     time.Time
	PreviousClose float64
	Change        float64
	ChangePercent string
}

// parseQuoteData will parse csv data from a reader
func parseQuoteData(r io.Reader) (*QuoteValue, error) {

	reader := csv.NewReader(r)

	// strip header
	if _, err := reader.Read(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	quoteValue := QuoteValue{}
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		value, err := parseQuoteRecord(record)
		if err != nil {
			return nil, err
		}
		quoteValue = *value
	}
	return &quoteValue, nil
}

// parseQuoteRecord will parse an individual csv record
func parseQuoteRecord(s []string) (*QuoteValue, error) {
	// these are the expected columns in the csv record
	const (
		symbol = iota
		open
		high
		low
		price
		volume
		latestDay
		previousClose
		change
		changePercent
	)

	value := &QuoteValue{}

	value.Symbol = s[symbol]

	f, err := parseFloat(s[open])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing open %s", s[open])
	}
	value.Open = f

	f, err = parseFloat(s[high])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing high %s", s[high])
	}
	value.High = f

	f, err = parseFloat(s[low])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing low %s", s[low])
	}
	value.Low = f

	f, err = parseFloat(s[price])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing price %s", s[price])
	}
	value.Price = f

	i, err := parseInt(s[volume])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing volume %s", s[volume])
	}
	value.Volume = i

	d, err := parseDate(s[latestDay], timeSeriesDateFormats...)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing latest date %s", s[latestDay])
	}
	value.LatestDay = d

	f, err = parseFloat(s[previousClose])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing previous close %s", s[previousClose])
	}
	value.PreviousClose = f

	f, err = parseFloat(s[change])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing change %s", s[change])
	}
	value.Change = f

	value.ChangePercent = s[changePercent]

	return value, nil
}
