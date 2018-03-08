package av

import (
	"testing"
)

func BenchmarkParseTimeSeriesData(b *testing.B) {
	buff := NewBuffCloser(sampleTimeSeriesData)
	for i := 0; i < b.N; i++ {
		buff.Restart()
		if _, err := parseTimeSeriesData(buff); err != nil {
			b.Fatalf("error parsing series: %v", err)
		}
	}
}
