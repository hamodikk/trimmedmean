package trimmedmean

import (
	"math"
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	data := []interface{}{14, 15.5, 15, 17, 22, 23, 23, 24, 25, 25.62, 26, 30, 31.8, 31, 32, 33, 34.3, 36, 38, 41}
	trim := 0.05
	expected := 26.79
	result, _ := TrimmedMean(data, trim)
	if math.Abs(result-expected) > 1e-6 {
		t.Errorf("Expected mean %f, got %f", expected, result)
	}
}
