package trimmedmean

import (
	"fmt"
	"sort"
)

// helper function to calculate the mean
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// function to sort and trim the dataset and find its mean
func TrimmedMean(data []interface{}, trims ...float64) (float64, error) {
	// error handling for empty dataset
	if len(data) == 0 {
		return 0, fmt.Errorf("data slice is empty")
	}

	// define the trim values based on arguments
	var lowerTrim, upperTrim float64
	switch len(trims) {
	case 1:
		lowerTrim = trims[0]
		upperTrim = trims[0]
	case 2:
		lowerTrim = trims[0]
		upperTrim = trims[1]
	default:
		return 0, fmt.Errorf("invalid number of trim values")
	}

	// error handling for trim value
	if lowerTrim < 0 || upperTrim < 0 || lowerTrim+upperTrim >= 1 {
		return 0, fmt.Errorf("trim value must be between 0 and 0.5")
	}

	// type assertion before sorting (+ error handling for non float/int values)
	floatData := make([]float64, len(data))
	for i, v := range data {
		switch v := v.(type) {
		case float64:
			floatData[i] = v
		case int:
			floatData[i] = float64(v)
		default:
			return 0, fmt.Errorf("data slice contains non-numeric values")
		}
	}

	// sort the data
	sort.Float64s(floatData)

	// trim the data
	lowerTrimAmount := int(float64(len(floatData)) * (lowerTrim))
	upperTrimAmount := int(float64(len(floatData)) * (upperTrim))
	trimmedData := floatData[lowerTrimAmount : len(floatData)-upperTrimAmount]

	// calculate the mean of the trimmed data
	return float64(mean(trimmedData)), nil
}
