package main

import (
	"errors"
	"math"
)

// DivideResult represents the result of a division operation
type DivideResult struct {
	Result   float64 `json:"result"`
	Overflow bool    `json:"overflow,omitempty"`
}

// DivideArrayResult represents the result of array division
type DivideArrayResult struct {
	Results []float64 `json:"results"`
	Error   string    `json:"error,omitempty"`
}

// BasicDivide performs basic division of two numbers
func BasicDivide(a, b float64) (DivideResult, error) {
	if b == 0 {
		return DivideResult{}, errors.New("division by zero")
	}
	
	result := a / b
	overflow := math.IsInf(result, 0) || math.IsNaN(result)
	
	return DivideResult{
		Result:   result,
		Overflow: overflow,
	}, nil
}

// DivideIntegers performs integer division with remainder
func DivideIntegers(a, b int64) (quotient int64, remainder int64, err error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	
	quotient = a / b
	remainder = a % b
	
	return quotient, remainder, nil
}

// DivideArray divides all elements in a slice by a divisor
func DivideArray(numbers []float64, divisor float64) (DivideArrayResult, error) {
	if divisor == 0 {
		return DivideArrayResult{}, errors.New("division by zero")
	}
	
	if len(numbers) == 0 {
		return DivideArrayResult{Results: []float64{}}, nil
	}
	
	results := make([]float64, len(numbers))
	
	for i, num := range numbers {
		results[i] = num / divisor
	}
	
	return DivideArrayResult{Results: results}, nil
}

// DivideArrayPairwise divides corresponding elements of two arrays
func DivideArrayPairwise(arr1, arr2 []float64) (DivideArrayResult, error) {
	if len(arr1) != len(arr2) {
		return DivideArrayResult{}, errors.New("arrays must have the same length")
	}
	
	if len(arr1) == 0 {
		return DivideArrayResult{Results: []float64{}}, nil
	}
	
	results := make([]float64, len(arr1))
	
	for i := 0; i < len(arr1); i++ {
		if arr2[i] == 0 {
			return DivideArrayResult{}, errors.New("division by zero at index " + string(rune(i+'0')))
		}
		results[i] = arr1[i] / arr2[i]
	}
	
	return DivideArrayResult{Results: results}, nil
}

// SafeDivide performs division with default value on division by zero
func SafeDivide(a, b, defaultValue float64) float64 {
	if b == 0 {
		return defaultValue
	}
	return a / b
}

// Modulo calculates the remainder of division
func Modulo(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("modulo by zero")
	}
	
	return math.Mod(a, b), nil
}

// Reciprocal calculates the reciprocal (1/x) of a number
func Reciprocal(x float64) (float64, error) {
	if x == 0 {
		return 0, errors.New("reciprocal of zero is undefined")
	}
	
	result := 1.0 / x
	if math.IsInf(result, 0) {
		return 0, errors.New("reciprocal overflow")
	}
	
	return result, nil
}
