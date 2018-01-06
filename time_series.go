package av

import (
	"encoding/csv"
	"io"
	"sort"
	"time"

	"github.com/pkg/errors"
)

type TimeSeries uint8

const (
	TimeSeriesDaily TimeSeries = iota
	TimeSeriesDailyAdjusted
	TimeSeriesWeekly
	TimeSeriesWeeklyAdjusted
	TimeSeriesMonthly
	TimeSeriesMonthlyAdjusted
	timeSeriesIntraday
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

func (t TimeSeries) KeyName() string {
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

func (t TimeInterval) KeyName() string {
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
	timeSeriesDateFormats = []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
	}
)

type TimeSeriesValue struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

// sortTimeSeriesValuesByDate allows DailyValues slices to be sorted by date in ascending order
type sortTimeSeriesValuesByDate []*TimeSeriesValue

func (b sortTimeSeriesValuesByDate) Len() int           { return len(b) }
func (b sortTimeSeriesValuesByDate) Less(i, j int) bool { return b[i].Date.Before(b[j].Date) }
func (b sortTimeSeriesValuesByDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

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

func parseTimeSeriesRecord(s []string) (*TimeSeriesValue, error) {
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
	value.Date = d

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
