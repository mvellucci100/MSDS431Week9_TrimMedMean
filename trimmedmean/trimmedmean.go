package trimmedmean

import (
	"sort"
	"fmt"
)

// TrimmedMean calculates the trimmed mean of a slice of numbers.
// The degreeOfTrimming argument is a proportion (between 0 and 1)
// representing the fraction of elements to trim from each end.
// If one argument is provided, it trims symmetrically from both ends.
func TrimmedMean(data []float64, degreeOfTrimming float64) (float64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("data slice is empty")
	}

	// Sort the data
	sort.Float64s(data)

	// Calculate the number of elements to trim
	n := len(data)
	if degreeOfTrimming < 0 || degreeOfTrimming > 0.5 {
		return 0, fmt.Errorf("degree of trimming must be between 0 and 0.5")
	}
	trimCount := int(degreeOfTrimming * float64(n))

	// Create a slice with the trimmed data
	trimmedData := data[trimCount : n-trimCount]

	// Compute the mean of the remaining data
	sum := 0.0
	for _, value := range trimmedData {
		sum += value
	}
	mean := sum / float64(len(trimmedData))

	return mean, nil
}
