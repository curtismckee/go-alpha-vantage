package av

import (
	"strconv"
	"time"
)

// parseFloat parses a float value.
// An error is returned if the value is not a float value.
func parseFloat(val string) (float64, error) {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// parseInt parses an int value.
// An error is returned if the value is not an int value.
func parseInt(val string) (int, error) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// parseDate parses a date value from a string.
// An error is returned if the value is not in the dateFormat format.
func parseDate(dateFormat string, v string) (time.Time, error) {
	t, err := time.Parse(dateFormat, v)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
