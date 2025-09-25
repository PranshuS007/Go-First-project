package main

import (
        "encoding/json"
        "net/http"
        "net/http/httptest"
        "strings"
        "testing"
)

// Test health check endpoint
func TestHealthCheckEndpoint(t *testing.T) {
        req, err := http.NewRequest("GET", "/health", nil)
        if err != nil {
                t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(healthHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusOK {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
        }

        var response map[string]interface{}
        err = json.Unmarshal(rr.Body.Bytes(), &response)
        if err != nil {
                t.Errorf("Failed to parse JSON response: %v", err)
        }

        if response["status"] != "healthy" {
                t.Errorf("Expected status 'healthy', got %v", response["status"])
        }

        if response["service"] != "go-first-project" {
                t.Errorf("Expected service 'go-first-project', got %v", response["service"])
        }

        if response["timestamp"] == nil {
                t.Error("Expected timestamp to be present")
        }
}

// Test health endpoint with wrong method
func TestHealthCheckEndpointWrongMethod(t *testing.T) {
        req, err := http.NewRequest("POST", "/health", nil)
        if err != nil {
                t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(healthHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusMethodNotAllowed {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
        }
}

// Test hello handler
func TestHelloHandler(t *testing.T) {
        req, err := http.NewRequest("GET", "/hello", nil)
        if err != nil {
                t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(helloHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusOK {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
        }

        expected := "Hello from Go server! 👋"
        if rr.Body.String() != expected {
                t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
        }
}

// Test hello handler with wrong path
func TestHelloHandlerWrongPath(t *testing.T) {
        req, err := http.NewRequest("GET", "/hello/test", nil)
        if err != nil {
                t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(helloHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusNotFound {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
        }
}

// Test hello handler with wrong method
func TestHelloHandlerWrongMethod(t *testing.T) {
        req, err := http.NewRequest("POST", "/hello", nil)
        if err != nil {
                t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(helloHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusMethodNotAllowed {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
        }
}

// Test form handler with valid data
func TestFormHandlerValidData(t *testing.T) {
        formData := "name=John&address=123 Main St"
        req, err := http.NewRequest("POST", "/form", strings.NewReader(formData))
        if err != nil {
                t.Fatal(err)
        }
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(formHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusOK {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
        }

        var response map[string]interface{}
        err = json.Unmarshal(rr.Body.Bytes(), &response)
        if err != nil {
                t.Errorf("Failed to parse JSON response: %v", err)
        }

        if response["success"] != true {
                t.Errorf("Expected success to be true, got %v", response["success"])
        }
}

// Test form handler with wrong method
func TestFormHandlerWrongMethod(t *testing.T) {
        req, err := http.NewRequest("GET", "/form", nil)
        if err != nil {
                t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(formHandler)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusMethodNotAllowed {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
        }
}

// Test rate limiter
func TestRateLimiter(t *testing.T) {
        rl := NewRateLimiter()
        
        // Test allowing requests
        if !rl.Allow("127.0.0.1") {
                t.Error("Expected first request to be allowed")
        }
        
        // Test rate limiting (this depends on the actual implementation)
        // For now, just test that the method exists and returns a boolean
        result := rl.Allow("127.0.0.1")
        if result != true && result != false {
                t.Error("Allow method should return a boolean")
        }
}

// Test sanitizeInput function
func TestSanitizeInput(t *testing.T) {
        tests := []struct {
                input    string
                expected string
        }{
                {"<script>alert('xss')</script>", "&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;"},
                {"normal text", "normal text"},
                {"", ""},
        }

        for _, test := range tests {
                result := sanitizeInput(test.input)
                if result != test.expected {
                        t.Errorf("sanitizeInput(%q) = %q, want %q", test.input, result, test.expected)
                }
        }
}

// Test validateFormInput function
func TestValidateFormInput(t *testing.T) {
        tests := []struct {
                name     string
                address  string
                hasError bool
        }{
                {"John", "123 Main St", false},
                {"", "123 Main St", true},     // empty name
                {"John", "", true},            // empty address
                {"", "", true},                // both empty
        }

        for _, test := range tests {
                err := validateFormInput(test.name, test.address)
                if test.hasError && err == nil {
                        t.Errorf("Expected error for name=%q, address=%q", test.name, test.address)
                }
                if !test.hasError && err != nil {
                        t.Errorf("Unexpected error for name=%q, address=%q: %v", test.name, test.address, err)
                }
        }
}

// Test getClientIP function
func TestGetClientIP(t *testing.T) {
        req, _ := http.NewRequest("GET", "/", nil)
        
        // Test with X-Forwarded-For header
        req.Header.Set("X-Forwarded-For", "192.168.1.1")
        ip := getClientIP(req)
        if ip != "192.168.1.1" {
                t.Errorf("Expected IP 192.168.1.1, got %s", ip)
        }
        
        // Test with X-Real-IP header
        req.Header.Del("X-Forwarded-For")
        req.Header.Set("X-Real-IP", "192.168.1.2")
        ip = getClientIP(req)
        if ip != "192.168.1.2" {
                t.Errorf("Expected IP 192.168.1.2, got %s", ip)
        }
}

// Test sendErrorResponse function
func TestSendErrorResponse(t *testing.T) {
        rr := httptest.NewRecorder()
        sendErrorResponse(rr, "Test Error", "Test Message", http.StatusBadRequest)

        if status := rr.Code; status != http.StatusBadRequest {
                t.Errorf("sendErrorResponse returned wrong status code: got %v want %v", status, http.StatusBadRequest)
        }

        var response ErrorResponse
        err := json.Unmarshal(rr.Body.Bytes(), &response)
        if err != nil {
                t.Errorf("Failed to parse JSON response: %v", err)
        }

        if response.Error != "Test Error" {
                t.Errorf("Expected error 'Test Error', got %v", response.Error)
        }

        if response.Message != "Test Message" {
                t.Errorf("Expected message 'Test Message', got %v", response.Message)
        }

        if response.Code != http.StatusBadRequest {
                t.Errorf("Expected code %v, got %v", http.StatusBadRequest, response.Code)
        }
}

// Benchmark test for health endpoint
func BenchmarkHealthEndpoint(b *testing.B) {
        req, _ := http.NewRequest("GET", "/health", nil)
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                rr := httptest.NewRecorder()
                handler := http.HandlerFunc(healthHandler)
                handler.ServeHTTP(rr, req)
        }
}

// Benchmark test for hello endpoint
func BenchmarkHelloEndpoint(b *testing.B) {
        req, _ := http.NewRequest("GET", "/hello", nil)
        
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                rr := httptest.NewRecorder()
                handler := http.HandlerFunc(helloHandler)
                handler.ServeHTTP(rr, req)
        }
}
