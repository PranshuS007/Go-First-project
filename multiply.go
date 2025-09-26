package main

import (
	"errors"
	"math"
)

// MultiplyResult represents the result of a multiplication operation
type MultiplyResult struct {
	Result   float64 `json:"result"`
	Overflow bool    `json:"overflow,omitempty"`
}

// MultiplyArrayResult represents the result of array multiplication
type MultiplyArrayResult struct {
	Results  []float64 `json:"results"`
	Overflow bool      `json:"overflow,omitempty"`
}

// BasicMultiply performs basic multiplication of two numbers
func BasicMultiply(a, b float64) MultiplyResult {
	result := a * b
	overflow := math.IsInf(result, 0) || math.IsNaN(result)
	
	return MultiplyResult{
		Result:   result,
		Overflow: overflow,
	}
}

// MultiplyIntegers performs integer multiplication with overflow detection
func MultiplyIntegers(a, b int64) (int64, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}
	
	// Check for overflow
	if a > 0 && b > 0 && a > math.MaxInt64/b {
		return 0, errors.New("integer overflow: result too large")
	}
	if a < 0 && b < 0 && a < math.MaxInt64/b {
		return 0, errors.New("integer overflow: result too large")
	}
	if (a > 0 && b < 0 && b < math.MinInt64/a) || (a < 0 && b > 0 && a < math.MinInt64/b) {
		return 0, errors.New("integer underflow: result too small")
	}
	
	return a * b, nil
}

// MultiplyArray multiplies all elements in a slice
func MultiplyArray(numbers []float64) MultiplyArrayResult {
	if len(numbers) == 0 {
		return MultiplyArrayResult{Results: []float64{}}
	}
	
	result := 1.0
	overflow := false
	
	for _, num := range numbers {
		result *= num
		if math.IsInf(result, 0) || math.IsNaN(result) {
			overflow = true
			break
		}
	}
	
	return MultiplyArrayResult{
		Results:  []float64{result},
		Overflow: overflow,
	}
}

// MultiplyArrayPairwise multiplies corresponding elements of two arrays
func MultiplyArrayPairwise(arr1, arr2 []float64) (MultiplyArrayResult, error) {
	if len(arr1) != len(arr2) {
		return MultiplyArrayResult{}, errors.New("arrays must have the same length")
	}
	
	if len(arr1) == 0 {
		return MultiplyArrayResult{Results: []float64{}}, nil
	}
	
	results := make([]float64, len(arr1))
	overflow := false
	
	for i := 0; i < len(arr1); i++ {
		results[i] = arr1[i] * arr2[i]
		if math.IsInf(results[i], 0) || math.IsNaN(results[i]) {
			overflow = true
		}
	}
	
	return MultiplyArrayResult{
		Results:  results,
		Overflow: overflow,
	}, nil
}

// MultiplyByScalar multiplies each element in an array by a scalar value
func MultiplyByScalar(numbers []float64, scalar float64) MultiplyArrayResult {
	if len(numbers) == 0 {
		return MultiplyArrayResult{Results: []float64{}}
	}
	
	results := make([]float64, len(numbers))
	overflow := false
	
	for i, num := range numbers {
		results[i] = num * scalar
		if math.IsInf(results[i], 0) || math.IsNaN(results[i]) {
			overflow = true
		}
	}
	
	return MultiplyArrayResult{
		Results:  results,
		Overflow: overflow,
	}
}

// Power calculates base raised to the power of exponent
func Power(base, exponent float64) MultiplyResult {
	result := math.Pow(base, exponent)
	overflow := math.IsInf(result, 0) || math.IsNaN(result)
	
	return MultiplyResult{
		Result:   result,
		Overflow: overflow,
	}
}

// Factorial calculates the factorial of a non-negative integer
func Factorial(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	
	if n == 0 || n == 1 {
		return 1, nil
	}
	
	// Check if factorial will overflow
	if n > 20 {
		return 0, errors.New("factorial overflow: number too large")
	}
	
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	
	return result, nil
}
