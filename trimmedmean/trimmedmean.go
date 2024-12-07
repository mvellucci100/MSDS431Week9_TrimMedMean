package trimmedmean

import (
	"fmt"
	"sort"
)

// TrimmedMean calculates the trimmed mean for both integer and float slices.
// degreeOfTrimming is the fraction (0 to 0.5) of the elements to trim from both ends.
func TrimmedMean(data interface{}, degreeOfTrimming float64) (float64, error) {
	if degreeOfTrimming < 0 || degreeOfTrimming > 0.5 {
		return 0, fmt.Errorf("degree of trimming must be between 0 and 0.5")
	}

	// Initialize an empty slice to store the data as float64
	var floatData []float64

	// Convert the input data into a slice of float64
	switch v := data.(type) {
	case []int:
		for _, val := range v {
			floatData = append(floatData, float64(val))
		}
	case []float64:
		floatData = append(floatData, v...)
	default:
		return 0, fmt.Errorf("unsupported data type: expected []int or []float64")
	}

	// Sort the data
	sort.Float64s(floatData)

	// Calculate the number of elements to trim
	n := len(floatData)
	trimCount := int(degreeOfTrimming * float64(n))

	// Create a slice with the trimmed data
	trimmedData := floatData[trimCount : n-trimCount]

	// Compute the mean of the remaining data
	sum := 0.0
	for _, value := range trimmedData {
		sum += value
	}
	mean := sum / float64(len(trimmedData))

	return mean, nil
}
