package av

import (
	"time"
	"sort"
	"encoding/csv"
	"io"
)

type TimeSeries uint8

const (
	TimeSeriesIntraday        TimeSeries = iota
	TimeSeriesDaily
	TimeSeriesDailyAdjusted
	TimeSeriesWeekly
	TimeSeriesWeeklyAdjusted
	TimeSeriesMonthly
	TimeSeriesMonthlyAdjusted
)

func (t TimeSeries) String() string {
	switch t {
	case TimeSeriesIntraday:
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

const (
	timeSeriesDateFormat = "2006-01-02"
)

type TimeSeriesValue struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}

// sortTimeSeriesValuesByDate allows DailyValues slices to be sorted by date in ascending order
type sortTimeSeriesValuesByDate []*TimeSeriesValue

func (b sortTimeSeriesValuesByDate) Len() int           { return len(b) }
func (b sortTimeSeriesValuesByDate) Less(i, j int) bool { return b[i].Date.Before(b[j].Date) }
func (b sortTimeSeriesValuesByDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func parseTimeSeriesData(r io.Reader) ([]*TimeSeriesValue, error) {

	reader := csv.NewReader(r)
	reader.ReuseRecord = true // optimization

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

	d, err := parseDate(timeSeriesDateFormat, s[timestamp])
	if err != nil {
		return nil, err
	}
	value.Date = d

	f, err := parseFloat(s[open])
	if err != nil {
		return nil, err
	}
	value.Open = f

	f, err = parseFloat(s[high])
	if err != nil {
		return nil, err
	}
	value.High = f

	f, err = parseFloat(s[low])
	if err != nil {
		return nil, err
	}
	value.Low = f

	f, err = parseFloat(s[close])
	if err != nil {
		return nil, err
	}
	value.Close = f

	i, err := parseInt(s[volume])
	if err != nil {
		return nil, err
	}
	value.Volume = i

	return value, nil
}
