package av

import (
	"fmt"
	"strconv"
	"time"
)

// parseFloat parses a float value from a map by key.
// An error is returned if the key does not exist or if the value is not a float value.
func parseFloat(m map[string]string, key string) (float64, error) {
	val, ok := m[key]
	if !ok {
		return 0, fmt.Errorf("key %s not found", key)
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// parseInt parses an int value from a map by key.
// An error is returned if the key does not exist or if the value is not an int value.
func parseInt(m map[string]string, key string) (int, error) {
	val, ok := m[key]
	if !ok {
		return 0, fmt.Errorf("key %s not found", key)
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// parseDate parses a date-only value from a map by key.
// An error is returned if the key does not exist or if the value is not in the dateFormat format.
func parseDate(dateFormat string, m map[string]string, key string) (time.Time, error) {
	val, ok := m[key]
	if !ok {
		return time.Time{}, fmt.Errorf("key %s not found", key)
	}
	t, err := time.Parse(dateFormat, val)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// parseString parses a string value from a map by key. an error is returned if the key does not exist
func parseString(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	return val, nil
}
