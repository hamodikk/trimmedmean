# Trimmed Mean Calculation Package

This package provides a function that calculates the trimmed mean of a provided slice in Go. Trimmed mean is a statistical measure that removes the outliers by a specific percentage before calculating the mean.

## Table of Contents
- [Features](#features)
- [Usage](#usage)
- [Testing](#testing)
- [Code Explanation](#code-explanation)

## Features

- Computes the trimmed mean of a slice of numbers.
- Accepts input for percentage of trimming.
- Works with symmetrical and asymmetrical trimming.
- Includes unit test and error handling for invalid input.

## Usage

To use the package, first clone the repository in your directory:
```bash
git clone https://github.com/hamodikk/trimmedmean.git
```
Change directory to the repository:
```bash
cd <path/to/repository/directory>
```
To use the package in your program, make sure to add the repository in your go.mod file:
```go
require (
    github.com/hamodikk/trimmedmean v1.0.0
)
```
Also make sure to import the package in your `main.go`.
```go
import (
    "github.com/hamodikk/trimmedmean"
```
Now, you can use the package in your code as follows:
```go
// asymmetrical trim
populationTrimmedMeanAsym, _ := trimmedmean.TrimmedMean(population, <lower_trim_value>, <upper_trim_value>)

// symmetrical trim
populationTrimmedMeanSym, _ := trimmedmean.TrimmedMean(population, <trim_value>)
```

## Testing

Here is how you can test the code:
```bash
# Test the code
go test -v
```

Example test output:
```bash
PASS
ok      path/to/trimmedmean 0.180s
```

## Code Explanation

The code consists of two functions. The first function is a helper function that calculates the mean of a dataset:
```go
// helper function to calculate the mean
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}
```

The other function is the `TrimmedMean` function that trims the dataset. This includes error handling, accepting single (symmetrical) or double (asymmetrical) trim values.

- Error handling for empty dataset
```go
    if len(data) == 0 {
		return 0, fmt.Errorf("data slice is empty")
	}
```

- Define trim values based on arguments. The function determines symmetry of trimming based on the number of trim inputs.
```go
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
```

- Error handling of trim value to make sure dataset is not overtrimmed.
```go
	if lowerTrim < 0 || upperTrim < 0 || lowerTrim+upperTrim >= 1 {
		return 0, fmt.Errorf("trim value must be between 0 and 0.5")
	}
```

- Type assertion. I do this to convert all the numbers to float64, in case the dataset has integers and floats.
```go
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
```

- Data sorting, followed by trimming.
```go
    // sort the data
	sort.Float64s(floatData)

	// trim the data
	lowerTrimAmount := int(float64(len(floatData)) * (lowerTrim))
	upperTrimAmount := int(float64(len(floatData)) * (upperTrim))
	trimmedData := floatData[lowerTrimAmount : len(floatData)-upperTrimAmount]
```

- Return the mean of the trimmed data.
```go
	return float64(mean(trimmedData)), nil
```