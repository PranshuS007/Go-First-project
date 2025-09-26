
# Go First Project 🚀

A simple, secure, and well-structured Go web server that demonstrates best practices for web development in Go. This project serves static files and handles form submissions with proper validation, security headers, rate limiting, and graceful shutdown.

## ✨ Features

- **Static File Serving**: Serves HTML, CSS, and other static assets
- **Form Handling**: Secure form processing with validation and sanitization
- **Security**: 
  - Security headers (XSS protection, content type options, frame options)
  - Input validation and sanitization
  - Rate limiting (100 requests per minute per IP)
- **Logging**: Structured request logging with timestamps
- **Configuration**: Environment-based configuration with sensible defaults
- **Graceful Shutdown**: Proper server shutdown handling
- **Responsive Design**: Mobile-friendly UI with modern CSS
- **Error Handling**: Consistent JSON error responses

## 🛠️ Installation

### Prerequisites

- Go 1.19 or higher
- Git

### Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/PranshuS007/Go-First-project.git
   cd Go-First-project
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables (optional)**
   ```bash
   cp .env.example .env
   # Edit .env file with your preferred settings
   ```

## 🚀 Usage

### Running the Server

1. **Default configuration (port 8080)**
   ```bash
   go run main.go
   ```

2. **Custom port using environment variable**
   ```bash
   PORT=3000 go run main.go
   ```

3. **Using environment file**
   ```bash
   # Set PORT=3000 in .env file
   go run main.go
   ```

### Building for Production

```bash
# Build the binary
go build -o server main.go

# Run the binary
./server
```

### Cross-platform Building

```bash
# For Linux
GOOS=linux GOARCH=amd64 go build -o server-linux main.go

# For Windows
GOOS=windows GOARCH=amd64 go build -o server-windows.exe main.go

# For macOS
GOOS=darwin GOARCH=amd64 go build -o server-macos main.go
```

## 📡 API Endpoints

### GET /
- **Description**: Serves the main static page
- **Response**: HTML page

### GET /hello
- **Description**: Simple hello endpoint
- **Response**: Plain text greeting
- **Example**:
  ```bash
  curl http://localhost:8080/hello
  # Response: Hello from Go server! 👋
  ```

### POST /form
- **Description**: Handles form submissions
- **Content-Type**: `application/x-www-form-urlencoded`
- **Parameters**:
  - `name` (required): User's name (max 100 characters)
  - `address` (required): User's address (max 200 characters)
- **Response**: JSON with success status and sanitized data
- **Example**:
  ```bash
  curl -X POST http://localhost:8080/form \
    -d "name=John Doe" \
    -d "address=123 Main St"
  ```

### GET /health
- **Description**: Health check endpoint for monitoring
- **Response**: JSON with service status and timestamp
- **Example**:
  ```bash
  curl http://localhost:8080/health
  # Response: {"status":"healthy","timestamp":1632150000,"service":"go-first-project","uptime":"1h0m0s"}
  ```

## 🧮 Multiplication API Endpoints

### POST /multiply
- **Description**: Basic multiplication of two numbers
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "a": 5.0,
    "b": 3.0
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "result": 15.0,
      "overflow": false
    }
  }
  ```

### POST /multiply/array
- **Description**: Multiply all elements in an array together
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "numbers": [2.0, 3.0, 4.0]
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "results": [24.0],
      "overflow": false
    }
  }
  ```

### POST /multiply/pairwise
- **Description**: Multiply corresponding elements of two arrays
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "array1": [2.0, 3.0, 4.0],
    "array2": [1.0, 2.0, 3.0]
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "results": [2.0, 6.0, 12.0],
      "overflow": false
    }
  }
  ```

### POST /multiply/scalar
- **Description**: Multiply each element in an array by a scalar value
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "numbers": [1.0, 2.0, 3.0],
    "scalar": 2.0
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "results": [2.0, 4.0, 6.0],
      "overflow": false
    }
  }
  ```

### POST /power
- **Description**: Calculate base raised to the power of exponent
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "base": 2.0,
    "exponent": 3.0
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "result": 8.0,
      "overflow": false
    }
  }
  ```

### POST /factorial
- **Description**: Calculate factorial of a non-negative integer (max 20)
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "number": 5
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "data": {
      "input": 5,
      "result": 120
    }
  }
  ```

### Error Responses

All endpoints return consistent JSON error responses:

```json
{
  "error": "Validation Error",
  "message": "Name is required",
  "code": 400
}
```

## 🏗️ Project Structure

```
Go-First-project/
├── main.go              # Main server code
├── go.mod              # Go module file
├── .env.example        # Environment variables template
├── .gitignore          # Git ignore rules
├── README.md           # This file
└── static/             # Static assets
    ├── index.html      # Home page
    ├── form.html       # Contact form page
    └── style.css       # Stylesheet
```

## 🔧 Configuration

The server can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `GO_ENV` | `development` | Environment mode |

## 🛡️ Security Features

- **Rate Limiting**: 100 requests per minute per IP address
- **Input Validation**: Server-side validation for all form inputs
- **Input Sanitization**: HTML escaping and length limits
- **Security Headers**: 
  - `X-Content-Type-Options: nosniff`
  - `X-Frame-Options: DENY`
  - `X-XSS-Protection: 1; mode=block`
  - `Referrer-Policy: strict-origin-when-cross-origin`
- **Method Validation**: Proper HTTP method checking
- **Error Handling**: No sensitive information in error responses

## 📝 Development

### Code Structure

- **main.go**: Contains all server logic including:
  - HTTP handlers (`helloHandler`, `formHandler`)
  - Middleware (logging, security headers, rate limiting)
  - Utility functions (validation, sanitization)
  - Graceful shutdown handling

### Adding New Features

1. **New Endpoints**: Add handler functions and register them in `main()`
2. **Middleware**: Create middleware functions and wrap them around handlers
3. **Static Assets**: Add files to the `static/` directory
4. **Configuration**: Add new environment variables and update `.env.example`

### Testing

```bash
# Run tests (when test files are added)
go test ./...

# Check code formatting
go fmt ./...

# Run static analysis
go vet ./...
```

## 🚀 Deployment

### Docker (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./server"]
```

Build and run:

```bash
docker build -t go-web-server .
docker run -p 8080:8080 go-web-server
```

### Production Considerations

- Use a reverse proxy (nginx, Apache) in production
- Set up proper logging aggregation
- Configure monitoring and health checks
- Use HTTPS with TLS certificates
- Set up database connections if needed
- Configure proper CORS headers for API endpoints

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- Built with Go's standard library
- Inspired by Go web development best practices
- Modern CSS design patterns
- Security recommendations from OWASP

## 📞 Support

If you have any questions or run into issues, please:

1. Check the existing issues on GitHub
2. Create a new issue with detailed information
3. Include Go version, OS, and error messages

---

**Happy coding! 🎉**
