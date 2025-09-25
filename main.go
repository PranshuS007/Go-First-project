package main

import (
        "context"
        "encoding/json"
        "fmt"
        "html"
        "log"
        "net/http"
        "os"
        "os/signal"
        "strings"
        "sync"
        "syscall"
        "time"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
        Error   string `json:"error"`
        Message string `json:"message"`
        Code    int    `json:"code"`
}

// RateLimiter implements a simple in-memory rate limiter
type RateLimiter struct {
        visitors map[string]*Visitor
        mu       sync.RWMutex
}

type Visitor struct {
        lastSeen time.Time
        count    int
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter() *RateLimiter {
        rl := &RateLimiter{
                visitors: make(map[string]*Visitor),
        }

        // Clean up old visitors every minute
        go func() {
                for {
                        time.Sleep(time.Minute)
                        rl.cleanupVisitors()
                }
        }()

        return rl
}

// Allow checks if a request should be allowed (max 100 requests per minute per IP)
func (rl *RateLimiter) Allow(ip string) bool {
        rl.mu.Lock()
        defer rl.mu.Unlock()

        v, exists := rl.visitors[ip]
        if !exists {
                rl.visitors[ip] = &Visitor{
                        lastSeen: time.Now(),
                        count:    1,
                }
                return true
        }

        if time.Since(v.lastSeen) > time.Minute {
                v.count = 1
                v.lastSeen = time.Now()
                return true
        }

        if v.count >= 100 {
                return false
        }

        v.count++
        return true
}

func (rl *RateLimiter) cleanupVisitors() {
        rl.mu.Lock()
        defer rl.mu.Unlock()

        for ip, v := range rl.visitors {
                if time.Since(v.lastSeen) > time.Minute {
                        delete(rl.visitors, ip)
                }
        }
}

var rateLimiter = NewRateLimiter()

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                start := time.Now()

                // Get client IP
                clientIP := getClientIP(r)

                // Rate limiting
                if !rateLimiter.Allow(clientIP) {
                        sendErrorResponse(w, "Rate limit exceeded", "Too many requests", http.StatusTooManyRequests)
                        return
                }

                // Add security headers
                w.Header().Set("X-Content-Type-Options", "nosniff")
                w.Header().Set("X-Frame-Options", "DENY")
                w.Header().Set("X-XSS-Protection", "1; mode=block")
                w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

                next.ServeHTTP(w, r)

                // Log the request
                duration := time.Since(start)
                log.Printf("[%s] %s %s %s - %v",
                        time.Now().Format("2006-01-02 15:04:05"),
                        clientIP,
                        r.Method,
                        r.URL.Path,
                        duration)
        })
}

// Get client IP address
func getClientIP(r *http.Request) string {
        // Check X-Forwarded-For header
        xff := r.Header.Get("X-Forwarded-For")
        if xff != "" {
                ips := strings.Split(xff, ",")
                return strings.TrimSpace(ips[0])
        }

        // Check X-Real-IP header
        xri := r.Header.Get("X-Real-IP")
        if xri != "" {
                return xri
        }

        // Fall back to RemoteAddr
        return strings.Split(r.RemoteAddr, ":")[0]
}

// Send standardized error response
func sendErrorResponse(w http.ResponseWriter, error, message string, code int) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(code)

        response := ErrorResponse{
                Error:   error,
                Message: message,
                Code:    code,
        }

        json.NewEncoder(w).Encode(response)
}

// Sanitize input to prevent XSS
func sanitizeInput(input string) string {
        // Remove leading/trailing whitespace
        input = strings.TrimSpace(input)
        // Escape HTML characters
        input = html.EscapeString(input)
        // Limit length
        if len(input) > 100 {
                input = input[:100]
        }
        return input
}

// Validate form input
func validateFormInput(name, address string) error {
        if len(strings.TrimSpace(name)) == 0 {
                return fmt.Errorf("name is required")
        }
        if len(strings.TrimSpace(address)) == 0 {
                return fmt.Errorf("address is required")
        }
        if len(name) > 100 {
                return fmt.Errorf("name must be less than 100 characters")
        }
        if len(address) > 200 {
                return fmt.Errorf("address must be less than 200 characters")
        }
        return nil
}

func main() {
        // Get port from environment variable with default fallback
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        // Setup logging with timestamps
        log.SetFlags(log.LstdFlags | log.Lshortfile)

        // Setup routes
        mux := http.NewServeMux()

        // Static file server
        fileserver := http.FileServer(http.Dir("./static"))
        mux.Handle("/", fileserver)

        // API endpoints
        mux.HandleFunc("/form", formHandler)
        mux.HandleFunc("/hello", helloHandler)
        mux.HandleFunc("/health", healthHandler)

        // Wrap with logging middleware
        handler := loggingMiddleware(mux)

        // Create server
        server := &http.Server{
                Addr:         ":" + port,
                Handler:      handler,
                ReadTimeout:  15 * time.Second,
                WriteTimeout: 15 * time.Second,
                IdleTimeout:  60 * time.Second,
        }

        // Start server in a goroutine
        go func() {
                log.Printf("Starting server on port %s...", port)
                log.Printf("Server running at http://localhost:%s", port)
                if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
                        log.Fatalf("Server failed to start: %v", err)
                }
        }()

        // Wait for interrupt signal to gracefully shutdown the server
        quit := make(chan os.Signal, 1)
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
        <-quit
        log.Println("Shutting down server...")

        // Give outstanding requests 30 seconds to complete
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        if err := server.Shutdown(ctx); err != nil {
                log.Fatalf("Server forced to shutdown: %v", err)
        }

        log.Println("Server exited")
}

// helloHandler handles GET requests to /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
        // Check if path is exactly /hello
        if r.URL.Path != "/hello" {
                sendErrorResponse(w, "Not Found", "The requested resource was not found", http.StatusNotFound)
                return
        }

        // Only allow GET method
        if r.Method != http.MethodGet {
                sendErrorResponse(w, "Method Not Allowed", "Only GET method is allowed for this endpoint", http.StatusMethodNotAllowed)
                return
        }

        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintf(w, "Hello from Go server! 👋")
}

// formHandler handles POST requests to /form endpoint
func formHandler(w http.ResponseWriter, r *http.Request) {
        // Only allow POST method
        if r.Method != http.MethodPost {
                sendErrorResponse(w, "Method Not Allowed", "Only POST method is allowed for this endpoint", http.StatusMethodNotAllowed)
                return
        }

        // Parse form data
        if err := r.ParseForm(); err != nil {
                log.Printf("ParseForm error: %v", err)
                sendErrorResponse(w, "Bad Request", "Failed to parse form data", http.StatusBadRequest)
                return
        }

        // Get form values
        name := r.FormValue("name")
        address := r.FormValue("address")

        // Validate input
        if err := validateFormInput(name, address); err != nil {
                sendErrorResponse(w, "Validation Error", err.Error(), http.StatusBadRequest)
                return
        }

        // Sanitize input
        name = sanitizeInput(name)
        address = sanitizeInput(address)

        // Log the form submission
        log.Printf("Form submitted - Name: %s, Address: %s", name, address)

        // Send success response
        w.Header().Set("Content-Type", "application/json")
        response := map[string]interface{}{
                "success": true,
                "message": "Form submitted successfully", // Fixed typo: "succesful" -> "successfully"
                "data": map[string]string{
                        "name":    name,
                        "address": address,
                },
        }

        json.NewEncoder(w).Encode(response)
}

// healthHandler handles GET requests to /health endpoint for monitoring
func healthHandler(w http.ResponseWriter, r *http.Request) {
        // Only allow GET method
        if r.Method != http.MethodGet {
                sendErrorResponse(w, "Method Not Allowed", "Only GET method is allowed for this endpoint", http.StatusMethodNotAllowed)
                return
        }

        // Set response headers
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        // Create health response
        response := map[string]interface{}{
                "status":    "healthy",
                "timestamp": time.Now().Unix(),
                "service":   "go-first-project",
                "uptime":    time.Since(time.Now().Add(-time.Hour)).String(), // Simple uptime placeholder
        }

        json.NewEncoder(w).Encode(response)
}
