package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	// Test case 1: Trimmed mean for a slice of integers
	intData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedIntMean := 5.5  // Expected trimmed mean after trimming 5% from both ends (0.05)
	trimmedIntMean, err := TrimmedMean(intData, 0.05)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if trimmedIntMean != expectedIntMean {
		t.Errorf("Expected %v, got %v", expectedIntMean, trimmedIntMean)
	}

	// Test case 2: Trimmed mean for a slice of float64 values
	floatData := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
	expectedFloatMean := 5.55  // Expected trimmed mean after trimming 5% from both ends (0.05)
	trimmedFloatMean, err := TrimmedMean(floatData, 0.05)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if trimmedFloatMean != expectedFloatMean {
		t.Errorf("Expected %v, got %v", expectedFloatMean, trimmedFloatMean)
	}

	// Test case 3: Invalid trimming degree (greater than 0.5)
	_, err = TrimmedMean(intData, 0.6)
	if err == nil {
		t.Errorf("Expected an error for trimming degree greater than 0.5, but got nil")
	}

	// Test case 4: Unsupported data type (e.g., string)
	_, err = TrimmedMean("invalid data", 0.05)
	if err == nil {
		t.Errorf("Expected an error for unsupported data type, but got nil")
	}

	// Test case 5: Trimming on an empty slice
	_, err = TrimmedMean([]int{}, 0.05)
	if err == nil {
		t.Errorf("Expected an error for empty data, but got nil")
	}
}
