package av

import (
	"strconv"
)

// Slice identifies one of the 24 past 30-day windows.
type Slice string

// AllSlices returns the past 24 slices from oldest to newest.
func AllSlices() []Slice {
	all := make([]Slice, 24)
	for i := 0; i < 24; i++ {
		all[23-i] = SliceAt(i)
	}
	return all
}

// SliceAt returns the slice for the given offset. offset must be between 0 and
// 23. An offset of 0 identifies the most recent 30-day window, an offset of 23
// the oldest.
func SliceAt(offset int) Slice {
	if offset < 0 || offset > 23 {
		panic("slice offset must be between 0 and 23")
	}
	pref := "year2"
	if offset > 11 {
		pref = "year1"
		offset = offset - 12
	}
	return Slice(pref + "month" + strconv.Itoa(12-offset))
}
