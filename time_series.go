package av

import (
	"encoding/csv"
	"io"
	"sort"
	"time"

	"github.com/pkg/errors"
)

// TimeSeries specifies a given time series to query for.
// For valid options, see the TimeSeries* package constants.
type TimeSeries uint8

const (
	TimeSeriesDaily TimeSeries = iota
	TimeSeriesDailyAdjusted
	TimeSeriesWeekly
	TimeSeriesWeeklyAdjusted
	TimeSeriesMonthly
	TimeSeriesMonthlyAdjusted
	timeSeriesIntraday // intentionally not exported
)

func (t TimeSeries) String() string {
	switch t {
	case timeSeriesIntraday:
		return "TimeSeriesIntraday"
	case TimeSeriesDaily:
		return "TimeSeriesDaily"
	case TimeSeriesDailyAdjusted:
		return "TimeSeriesDailyAdjusted"
	case TimeSeriesWeekly:
		return "TimeSeriesWeekly"
	case TimeSeriesWeeklyAdjusted:
		return "TimeSeriesWeeklyAdjusted"
	case TimeSeriesMonthly:
		return "TimeSeriesMonthly"
	case TimeSeriesMonthlyAdjusted:
		return "TimeSeriesMonthlyAdjusted"
	}
	return "TimeSeriesUnknown"
}

// keyName returns the name of the TimeSeries used for Alpha Vantage API
func (t TimeSeries) keyName() string {
	switch t {
	case timeSeriesIntraday:
		return "TIME_SERIES_INTRADAY"
	case TimeSeriesDaily:
		return "TIME_SERIES_DAILY"
	case TimeSeriesDailyAdjusted:
		return "TIME_SERIES_DAILY_ADJUSTED"
	case TimeSeriesWeekly:
		return "TIME_SERIES_WEEKLY"
	case TimeSeriesWeeklyAdjusted:
		return "TIME_SERIES_WEEKLY_ADJUSTED"
	case TimeSeriesMonthly:
		return "TIME_SERIES_MONTHLY"
	case TimeSeriesMonthlyAdjusted:
		return "TIME_SERIES_MONTHLY_ADJUSTED"
	}
	return "UNKNOWN"
}

// TimeInterval specifies a frequency to query for intraday stock data.
// For valid options, see the TimeInterval* package constants.
type TimeInterval uint8

const (
	TimeIntervalOneMinute TimeInterval = iota
	TimeIntervalFiveMinute
	TimeIntervalFifteenMinute
	TimeIntervalThirtyMinute
	TimeIntervalSixtyMinute
)

func (t TimeInterval) String() string {
	switch t {
	case TimeIntervalOneMinute:
		return "TimeIntervalOneMinute"
	case TimeIntervalFiveMinute:
		return "TimeIntervalFiveMinute"
	case TimeIntervalFifteenMinute:
		return "TimeIntervalFifteenMinute"
	case TimeIntervalThirtyMinute:
		return "TimeIntervalThirtyMinute"
	case TimeIntervalSixtyMinute:
		return "TimeIntervalSixtyMinute"
	}
	return "TimeIntervalUnknown"
}

// keyName returns the name of the TimeInterval used for Alpha Vantage API
func (t TimeInterval) keyName() string {
	switch t {
	case TimeIntervalOneMinute:
		return "1min"
	case TimeIntervalFiveMinute:
		return "5min"
	case TimeIntervalFifteenMinute:
		return "15min"
	case TimeIntervalThirtyMinute:
		return "30min"
	case TimeIntervalSixtyMinute:
		return "60min"
	}
	return "unknown"
}

var (
	// timeSeriesDateFormats are the expected date formats in time series data
	timeSeriesDateFormats = []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
	}
)

// TimeSeriesValue is a piece of data for a given time about stock prices
type TimeSeriesValue struct {
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

// sortTimeSeriesValuesByDate allows TimeSeriesValue
// slices to be sorted by date in ascending order
type sortTimeSeriesValuesByDate []*TimeSeriesValue

func (b sortTimeSeriesValuesByDate) Len() int           { return len(b) }
func (b sortTimeSeriesValuesByDate) Less(i, j int) bool { return b[i].Time.Before(b[j].Time) }
func (b sortTimeSeriesValuesByDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// parseTimeSeriesData will parse csv data from a reader
func parseTimeSeriesData(r io.Reader) ([]*TimeSeriesValue, error) {

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

	values := make([]*TimeSeriesValue, 0, 64)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		value, err := parseTimeSeriesRecord(record)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	// sort values by date
	sort.Sort(sortTimeSeriesValuesByDate(values))

	return values, nil

}

// parseDigitalCurrencySeriesRecord will parse an individual csv record
func parseTimeSeriesRecord(s []string) (*TimeSeriesValue, error) {
	// these are the expected columns in the csv record
	const (
		timestamp = iota
		open
		high
		low
		close
		volume
	)

	value := &TimeSeriesValue{}

	d, err := parseDate(s[timestamp], timeSeriesDateFormats...)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing timestamp %s", s[timestamp])
	}
	value.Time = d

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

	f, err = parseFloat(s[close])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing close %s", s[close])
	}
	value.Close = f

	f, err = parseFloat(s[volume])
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing volume %s", s[volume])
	}
	value.Volume = f

	return value, nil
}
