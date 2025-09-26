package main

import (
        "bytes"
        "encoding/json"
        "math"
        "net/http"
        "net/http/httptest"
        "testing"
)

// Test BasicMultiply function
func TestBasicMultiply(t *testing.T) {
        tests := []struct {
                name     string
                a, b     float64
                expected float64
                overflow bool
        }{
                {"positive numbers", 5.0, 3.0, 15.0, false},
                {"negative numbers", -4.0, -2.0, 8.0, false},
                {"mixed signs", -6.0, 3.0, -18.0, false},
                {"zero multiplication", 0.0, 100.0, 0.0, false},
                {"decimal numbers", 2.5, 4.0, 10.0, false},
                {"large numbers", 1e10, 1e10, 1e20, false},
                {"overflow case", math.MaxFloat64, 2.0, math.Inf(1), true},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := BasicMultiply(tt.a, tt.b)
                        if result.Overflow != tt.overflow {
                                t.Errorf("BasicMultiply(%v, %v) overflow = %v, want %v", tt.a, tt.b, result.Overflow, tt.overflow)
                        }
                        if !tt.overflow && result.Result != tt.expected {
                                t.Errorf("BasicMultiply(%v, %v) = %v, want %v", tt.a, tt.b, result.Result, tt.expected)
                        }
                })
        }
}

// Test MultiplyIntegers function
func TestMultiplyIntegers(t *testing.T) {
        tests := []struct {
                name        string
                a, b        int64
                expected    int64
                expectError bool
        }{
                {"positive integers", 5, 3, 15, false},
                {"negative integers", -4, -2, 8, false},
                {"mixed signs", -6, 3, -18, false},
                {"zero multiplication", 0, 100, 0, false},
                {"large numbers", 1000000, 1000000, 1000000000000, false},
                {"overflow case", math.MaxInt64, 2, 0, true},
                {"underflow case", math.MinInt64, 2, 0, true},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result, err := MultiplyIntegers(tt.a, tt.b)
                        if tt.expectError {
                                if err == nil {
                                        t.Errorf("MultiplyIntegers(%v, %v) expected error but got none", tt.a, tt.b)
                                }
                        } else {
                                if err != nil {
                                        t.Errorf("MultiplyIntegers(%v, %v) unexpected error: %v", tt.a, tt.b, err)
                                }
                                if result != tt.expected {
                                        t.Errorf("MultiplyIntegers(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
                                }
                        }
                })
        }
}

// Test MultiplyArray function
func TestMultiplyArray(t *testing.T) {
        tests := []struct {
                name     string
                numbers  []float64
                expected float64
                overflow bool
        }{
                {"empty array", []float64{}, 0, false},
                {"single element", []float64{5.0}, 5.0, false},
                {"positive numbers", []float64{2.0, 3.0, 4.0}, 24.0, false},
                {"with zero", []float64{2.0, 0.0, 4.0}, 0.0, false},
                {"negative numbers", []float64{-2.0, 3.0, -4.0}, 24.0, false},
                {"decimal numbers", []float64{1.5, 2.0, 3.0}, 9.0, false},
                {"large numbers causing overflow", []float64{1e200, 1e200, 1e200}, math.Inf(1), true},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := MultiplyArray(tt.numbers)
                        if len(tt.numbers) == 0 {
                                if len(result.Results) != 0 {
                                        t.Errorf("MultiplyArray(%v) expected empty results", tt.numbers)
                                }
                                return
                        }
                        if result.Overflow != tt.overflow {
                                t.Errorf("MultiplyArray(%v) overflow = %v, want %v", tt.numbers, result.Overflow, tt.overflow)
                        }
                        if !tt.overflow && len(result.Results) > 0 && result.Results[0] != tt.expected {
                                t.Errorf("MultiplyArray(%v) = %v, want %v", tt.numbers, result.Results[0], tt.expected)
                        }
                })
        }
}

// Test MultiplyArrayPairwise function
func TestMultiplyArrayPairwise(t *testing.T) {
        tests := []struct {
                name        string
                arr1, arr2  []float64
                expected    []float64
                expectError bool
                overflow    bool
        }{
                {"equal length arrays", []float64{2.0, 3.0, 4.0}, []float64{1.0, 2.0, 3.0}, []float64{2.0, 6.0, 12.0}, false, false},
                {"empty arrays", []float64{}, []float64{}, []float64{}, false, false},
                {"single element", []float64{5.0}, []float64{3.0}, []float64{15.0}, false, false},
                {"different lengths", []float64{1.0, 2.0}, []float64{1.0}, nil, true, false},
                {"with negatives", []float64{-2.0, 3.0}, []float64{4.0, -1.0}, []float64{-8.0, -3.0}, false, false},
                {"with zero", []float64{0.0, 5.0}, []float64{10.0, 2.0}, []float64{0.0, 10.0}, false, false},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result, err := MultiplyArrayPairwise(tt.arr1, tt.arr2)
                        if tt.expectError {
                                if err == nil {
                                        t.Errorf("MultiplyArrayPairwise(%v, %v) expected error but got none", tt.arr1, tt.arr2)
                                }
                                return
                        }
                        if err != nil {
                                t.Errorf("MultiplyArrayPairwise(%v, %v) unexpected error: %v", tt.arr1, tt.arr2, err)
                                return
                        }
                        if result.Overflow != tt.overflow {
                                t.Errorf("MultiplyArrayPairwise(%v, %v) overflow = %v, want %v", tt.arr1, tt.arr2, result.Overflow, tt.overflow)
                        }
                        if len(result.Results) != len(tt.expected) {
                                t.Errorf("MultiplyArrayPairwise(%v, %v) result length = %v, want %v", tt.arr1, tt.arr2, len(result.Results), len(tt.expected))
                                return
                        }
                        for i, expected := range tt.expected {
                                if result.Results[i] != expected {
                                        t.Errorf("MultiplyArrayPairwise(%v, %v) result[%d] = %v, want %v", tt.arr1, tt.arr2, i, result.Results[i], expected)
                                }
                        }
                })
        }
}

// Test MultiplyByScalar function
func TestMultiplyByScalar(t *testing.T) {
        tests := []struct {
                name     string
                numbers  []float64
                scalar   float64
                expected []float64
                overflow bool
        }{
                {"positive scalar", []float64{1.0, 2.0, 3.0}, 2.0, []float64{2.0, 4.0, 6.0}, false},
                {"negative scalar", []float64{1.0, -2.0, 3.0}, -2.0, []float64{-2.0, 4.0, -6.0}, false},
                {"zero scalar", []float64{1.0, 2.0, 3.0}, 0.0, []float64{0.0, 0.0, 0.0}, false},
                {"empty array", []float64{}, 5.0, []float64{}, false},
                {"decimal scalar", []float64{2.0, 4.0}, 1.5, []float64{3.0, 6.0}, false},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := MultiplyByScalar(tt.numbers, tt.scalar)
                        if result.Overflow != tt.overflow {
                                t.Errorf("MultiplyByScalar(%v, %v) overflow = %v, want %v", tt.numbers, tt.scalar, result.Overflow, tt.overflow)
                        }
                        if len(result.Results) != len(tt.expected) {
                                t.Errorf("MultiplyByScalar(%v, %v) result length = %v, want %v", tt.numbers, tt.scalar, len(result.Results), len(tt.expected))
                                return
                        }
                        for i, expected := range tt.expected {
                                if result.Results[i] != expected {
                                        t.Errorf("MultiplyByScalar(%v, %v) result[%d] = %v, want %v", tt.numbers, tt.scalar, i, result.Results[i], expected)
                                }
                        }
                })
        }
}

// Test Power function
func TestPower(t *testing.T) {
        tests := []struct {
                name     string
                base     float64
                exponent float64
                expected float64
                overflow bool
        }{
                {"positive base and exponent", 2.0, 3.0, 8.0, false},
                {"negative base, even exponent", -2.0, 2.0, 4.0, false},
                {"negative base, odd exponent", -2.0, 3.0, -8.0, false},
                {"zero base", 0.0, 5.0, 0.0, false},
                {"base to power of zero", 5.0, 0.0, 1.0, false},
                {"base to power of one", 5.0, 1.0, 5.0, false},
                {"fractional exponent", 4.0, 0.5, 2.0, false},
                {"large result", 2.0, 10.0, 1024.0, false},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := Power(tt.base, tt.exponent)
                        if result.Overflow != tt.overflow {
                                t.Errorf("Power(%v, %v) overflow = %v, want %v", tt.base, tt.exponent, result.Overflow, tt.overflow)
                        }
                        if !tt.overflow && math.Abs(result.Result-tt.expected) > 1e-10 {
                                t.Errorf("Power(%v, %v) = %v, want %v", tt.base, tt.exponent, result.Result, tt.expected)
                        }
                })
        }
}

// Test Factorial function
func TestFactorial(t *testing.T) {
        tests := []struct {
                name        string
                n           int
                expected    int64
                expectError bool
        }{
                {"factorial of 0", 0, 1, false},
                {"factorial of 1", 1, 1, false},
                {"factorial of 5", 5, 120, false},
                {"factorial of 10", 10, 3628800, false},
                {"factorial of 20", 20, 2432902008176640000, false},
                {"negative number", -1, 0, true},
                {"too large number", 25, 0, true},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result, err := Factorial(tt.n)
                        if tt.expectError {
                                if err == nil {
                                        t.Errorf("Factorial(%v) expected error but got none", tt.n)
                                }
                        } else {
                                if err != nil {
                                        t.Errorf("Factorial(%v) unexpected error: %v", tt.n, err)
                                }
                                if result != tt.expected {
                                        t.Errorf("Factorial(%v) = %v, want %v", tt.n, result, tt.expected)
                                }
                        }
                })
        }
}

// Test multiplyHandler endpoint
func TestMultiplyHandler(t *testing.T) {
        tests := []struct {
                name           string
                method         string
                path           string
                body           interface{}
                expectedStatus int
                expectSuccess  bool
        }{
                {
                        name:           "valid multiplication",
                        method:         "POST",
                        path:           "/multiply",
                        body:           MultiplyRequest{A: 5.0, B: 3.0},
                        expectedStatus: http.StatusOK,
                        expectSuccess:  true,
                },
                {
                        name:           "wrong method",
                        method:         "GET",
                        path:           "/multiply",
                        body:           nil,
                        expectedStatus: http.StatusMethodNotAllowed,
                        expectSuccess:  false,
                },
                {
                        name:           "wrong path",
                        method:         "POST",
                        path:           "/multiply/wrong",
                        body:           MultiplyRequest{A: 5.0, B: 3.0},
                        expectedStatus: http.StatusNotFound,
                        expectSuccess:  false,
                },
                {
                        name:           "invalid JSON",
                        method:         "POST",
                        path:           "/multiply",
                        body:           "invalid json",
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
                {
                        name:           "numbers too large",
                        method:         "POST",
                        path:           "/multiply",
                        body:           MultiplyRequest{A: 1e16, B: 3.0},
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        var body []byte
                        var err error
                        if tt.body != nil {
                                if str, ok := tt.body.(string); ok {
                                        body = []byte(str)
                                } else {
                                        body, err = json.Marshal(tt.body)
                                        if err != nil {
                                                t.Fatalf("Failed to marshal request body: %v", err)
                                        }
                                }
                        }

                        req := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(body))
                        req.Header.Set("Content-Type", "application/json")
                        w := httptest.NewRecorder()

                        multiplyHandler(w, req)

                        if w.Code != tt.expectedStatus {
                                t.Errorf("multiplyHandler() status = %v, want %v", w.Code, tt.expectedStatus)
                        }

                        if tt.expectSuccess {
                                var response map[string]interface{}
                                if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
                                        t.Errorf("Failed to unmarshal response: %v", err)
                                }
                                if success, ok := response["success"].(bool); !ok || !success {
                                        t.Errorf("Expected successful response, got: %v", response)
                                }
                        }
                })
        }
}

// Test multiplyArrayHandler endpoint
func TestMultiplyArrayHandler(t *testing.T) {
        tests := []struct {
                name           string
                method         string
                path           string
                body           interface{}
                expectedStatus int
                expectSuccess  bool
        }{
                {
                        name:           "valid array multiplication",
                        method:         "POST",
                        path:           "/multiply/array",
                        body:           ArrayRequest{Numbers: []float64{2.0, 3.0, 4.0}},
                        expectedStatus: http.StatusOK,
                        expectSuccess:  true,
                },
                {
                        name:           "empty array",
                        method:         "POST",
                        path:           "/multiply/array",
                        body:           ArrayRequest{Numbers: []float64{}},
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
                {
                        name:           "array too large",
                        method:         "POST",
                        path:           "/multiply/array",
                        body:           ArrayRequest{Numbers: make([]float64, 1001)},
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
                {
                        name:           "numbers too large",
                        method:         "POST",
                        path:           "/multiply/array",
                        body:           ArrayRequest{Numbers: []float64{1e11, 2.0}},
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        body, err := json.Marshal(tt.body)
                        if err != nil {
                                t.Fatalf("Failed to marshal request body: %v", err)
                        }

                        req := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(body))
                        req.Header.Set("Content-Type", "application/json")
                        w := httptest.NewRecorder()

                        multiplyArrayHandler(w, req)

                        if w.Code != tt.expectedStatus {
                                t.Errorf("multiplyArrayHandler() status = %v, want %v", w.Code, tt.expectedStatus)
                        }

                        if tt.expectSuccess {
                                var response map[string]interface{}
                                if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
                                        t.Errorf("Failed to unmarshal response: %v", err)
                                }
                                if success, ok := response["success"].(bool); !ok || !success {
                                        t.Errorf("Expected successful response, got: %v", response)
                                }
                        }
                })
        }
}

// Test factorialHandler endpoint
func TestFactorialHandler(t *testing.T) {
        tests := []struct {
                name           string
                method         string
                path           string
                body           interface{}
                expectedStatus int
                expectSuccess  bool
        }{
                {
                        name:           "valid factorial",
                        method:         "POST",
                        path:           "/factorial",
                        body:           FactorialRequest{Number: 5},
                        expectedStatus: http.StatusOK,
                        expectSuccess:  true,
                },
                {
                        name:           "negative number",
                        method:         "POST",
                        path:           "/factorial",
                        body:           FactorialRequest{Number: -1},
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
                {
                        name:           "number too large",
                        method:         "POST",
                        path:           "/factorial",
                        body:           FactorialRequest{Number: 25},
                        expectedStatus: http.StatusBadRequest,
                        expectSuccess:  false,
                },
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        body, err := json.Marshal(tt.body)
                        if err != nil {
                                t.Fatalf("Failed to marshal request body: %v", err)
                        }

                        req := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(body))
                        req.Header.Set("Content-Type", "application/json")
                        w := httptest.NewRecorder()

                        factorialHandler(w, req)

                        if w.Code != tt.expectedStatus {
                                t.Errorf("factorialHandler() status = %v, want %v", w.Code, tt.expectedStatus)
                        }

                        if tt.expectSuccess {
                                var response map[string]interface{}
                                if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
                                        t.Errorf("Failed to unmarshal response: %v", err)
                                }
                                if success, ok := response["success"].(bool); !ok || !success {
                                        t.Errorf("Expected successful response, got: %v", response)
                                }
                        }
                })
        }
}
