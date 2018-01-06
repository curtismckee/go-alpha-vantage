package av

import (
	"encoding/csv"
	"io"
	"sort"
	"time"

	"github.com/pkg/errors"
)

const (
	digitalCurrencySeriesDateFormat = "2006-01-02 15:04:05"
)

type DigitalCurrencySeriesValue struct {
	Date      time.Time
	Price     float64
	Volume    float64
	MarketCap float64
}

// sortDigitalCurrencySeriesValuesByDate allows DailyValues slices to be sorted by date in ascending order
type sortDigitalCurrencySeriesValuesByDate []*DigitalCurrencySeriesValue

func (b sortDigitalCurrencySeriesValuesByDate) Len() int           { return len(b) }
func (b sortDigitalCurrencySeriesValuesByDate) Less(i, j int) bool { return b[i].Date.Before(b[j].Date) }
func (b sortDigitalCurrencySeriesValuesByDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func parseDigitalCurrencySeriesData(r io.Reader) ([]*DigitalCurrencySeriesValue, error) {

	reader := csv.NewReader(r)
	reader.ReuseRecord = true // optimization
	reader.LazyQuotes = true
	reader.TrailingComma = true
	reader.TrimLeadingSpace = true

	// strip header
	if _, err := reader.Read(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	values := make([]*DigitalCurrencySeriesValue, 0, 64)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		value, err := parseDigitalCurrencySeriesRecord(record)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	// sort values by date
	sort.Sort(sortDigitalCurrencySeriesValuesByDate(values))

	return values, nil

}

func parseDigitalCurrencySeriesRecord(s []string) (*DigitalCurrencySeriesValue, error) {
	// [timestamp price (USD) price (USD) volume market cap (USD)]
	const (
		timestamp = iota
		price
		_
		volume
		marketCap
	)

	value := &DigitalCurrencySeriesValue{}

	d, err := parseDate(s[timestamp], digitalCurrencySeriesDateFormat)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing timestamp %s", s[timestamp])
	}
	value.Date = d

	f, err := parseFloat(s[price])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing price %s", s[price])
	}
	value.Price = f

	f, err = parseFloat(s[volume])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing volume %s", s[volume])
	}
	value.Volume = f

	f, err = parseFloat(s[marketCap])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing market cap %s", s[marketCap])
	}
	value.MarketCap = f

	return value, nil
}
