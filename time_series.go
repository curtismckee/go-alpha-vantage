package av

import (
	"time"
	"encoding/json"
	"net/http"
	"sort"
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

type TimeSeriesData struct {
	Meta   *TimeSeriesMeta
	Values []*TimeSeriesValues
}

type TimeSeriesMeta struct {
	Information   string
	Symbol        string
	LastRefreshed time.Time
	OutputSize    string
	TimeZone      string
}

type TimeSeriesValues struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}

// sortTimeSeriesValuesByDate allows DailyValues slices to be sorted by date in ascending order
type sortTimeSeriesValuesByDate []*TimeSeriesValues

func (b sortTimeSeriesValuesByDate) Len() int           { return len(b) }
func (b sortTimeSeriesValuesByDate) Less(i, j int) bool { return b[i].Date.Before(b[j].Date) }
func (b sortTimeSeriesValuesByDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func parseTimeSeriesData(res *http.Response) (*TimeSeriesData, error) {
	defer res.Body.Close()
	timeSeries := &TimeSeriesData{}
	if err := json.NewDecoder(res.Body).Decode(timeSeries); err != nil {
		return nil, err
	}

	// sort dates
	sort.Sort(sortTimeSeriesValuesByDate(timeSeries.Values))

	return timeSeries, nil
}

func (m *TimeSeriesMeta) UnmarshalJSON(b []byte) error {
	values := make(map[string]string)
	if err := json.Unmarshal(b, &values); err != nil {
		return err
	}

	s, err := parseString(values, "1. Information")
	if err != nil {
		return err
	}
	m.Information = s

	s, err = parseString(values, "2. Symbol")
	if err != nil {
		return err
	}
	m.Symbol = s

	d, err := parseDate(timeSeriesDateFormat, values, "3. Last Refreshed")
	if err != nil {
		return err
	}
	m.LastRefreshed = d

	s, err = parseString(values, "4. Output Size")
	if err != nil {
		return err
	}
	m.OutputSize = s

	s, err = parseString(values, "5. Time Zone")
	if err != nil {
		return err
	}
	m.TimeZone = s

	return nil
}

func (s *TimeSeriesValues) UnmarshalJSON(b []byte) error {
	values := make(map[string]string)
	if err := json.Unmarshal(b, &values); err != nil {
		return err
	}

	f, err := parseFloat(values, "1. open")
	if err != nil {
		return err
	}
	s.Open = f

	f, err = parseFloat(values, "2. high")
	if err != nil {
		return err
	}
	s.High = f

	f, err = parseFloat(values, "3. low")
	if err != nil {
		return err
	}
	s.Low = f

	f, err = parseFloat(values, "4. close")
	if err != nil {
		return err
	}
	s.Close = f

	i, err := parseInt(values, "5. volume")
	if err != nil {
		return err
	}
	s.Volume = i

	return nil
}
