# MSDS431Week9_TrimMedMean

# Unit Test for `TrimmedMean` Function

The `TrimmedMean` function is designed to calculate the trimmed mean for both integer and float slices. To ensure its correctness, we wrote a set of unit tests to verify its behavior under different scenarios. Below is an explanation of each test case used in the unit tests.

### Test Case 1: Trimmed Mean for Integers
In this test, we provide a slice of integers and check if the function correctly calculates the trimmed mean. The expected trimmed mean is calculated after trimming 5% of the elements from both ends (i.e., a trimming degree of 0.05). This ensures that the function works as expected with integer slices.

### Test Case 2: Trimmed Mean for Floats
This test checks if the function correctly computes the trimmed mean when provided with a slice of float64 values. Similar to the previous test, 5% of the elements are trimmed from both ends. The expected result is again compared to ensure correctness.

### Test Case 3: Invalid Trimming degree (Greater than 0.5)
This test verifies that the function properly handles an invalid trimming degree. According to the specification, the degree of trimming must be between 0 and 0.5. If a value greater than 0.5 is provided, the function should return an error.

### Test Case 4: Unsupported Data Type
This test checks if the function can handle unsupported data types. In this case, a string is passed instead of an integer or float slice. The function should return an error indicating that the data type is unsupported.

### Test Case 5: Empty Slice
In this test, we provide an empty slice of integers and expect the function to return an error. This ensures that the function handles empty input correctly and returns an appropriate error.
