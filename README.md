```markdown
# Go-First-project

A simple webserver written in Go, featuring basic arithmetic APIs, a custom linked list implementation, static file serving, and a simple in-memory rate limiter.

---

## Table of Contents

- [Description](#description)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Key Features](#key-features)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [License](#license)

---

## Description

**Go-First-project** is a beginner-friendly Go webserver that demonstrates:

- RESTful endpoints for arithmetic operations (multiplication, division)
- Static file serving (HTML/CSS)
- A custom singly linked list implementation
- Basic in-memory rate limiting per IP
- Health check and hello endpoints
- Example unit tests for handlers and core logic

---

## Technologies Used

- **Language:** Go (Golang)
- **Web:** net/http
- **Testing:** Go's built-in testing package
- **Other:** HTML, CSS (for static files)

---

## Project Structure

```
.
├── .env.example                # Example environment variables (if needed)
├── .gitignore
├── README.md
├── datastructures/
│   └── linkedlist.go           # Custom singly linked list implementation
├── divide.go                   # Division logic and helpers
├── go.mod                      # Go module definition
├── main.go                     # Main webserver, routing, rate limiting
├── main_test.go                # Tests for web handlers
├── multiply.go                 # Multiplication logic and helpers
├── multiply_test.go            # Tests for multiplication logic
└── static/
    ├── form.html               # Example HTML form
    ├── index.html              # Main static page
    └── style.css               # Stylesheet for static pages
```

---

## Key Features and Components

- **Webserver (main.go):**
  - Serves static files from `/static`
  - REST endpoints for arithmetic operations
  - `/health` endpoint for health checks
  - `/hello` endpoint for a simple greeting
  - In-memory rate limiter (100 requests/minute per IP)

- **Arithmetic APIs:**
  - **Multiplication** (`multiply.go`): Basic, integer, array, and pairwise multiplication with overflow detection
  - **Division** (`divide.go`): Basic, integer, array, and pairwise division with error handling

- **Data Structures:**
  - **Linked List** (`datastructures/linkedlist.go`): Custom singly linked list with add, delete, traverse, and search methods

- **Static Files:**
  - HTML and CSS files for basic web UI

- **Testing:**
  - Unit tests for handlers and arithmetic logic

---

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/PranshuS007/Go-First-project.git
   cd Go-First-project
   ```

2. **(Optional) Set up environment variables:**
   - Copy `.env.example` to `.env` and edit as needed.

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Run the server:**
   ```sh
   go run main.go
   ```
   The server will start (by default on port 8080 unless otherwise specified in the code).

---

## Usage

### Web Endpoints

- **Health Check**
  - `GET /health`
  - Returns JSON with service status and timestamp.

- **Hello**
  - `GET /hello`
  - Returns a simple greeting.

- **Static Files**
  - Access via `/static/index.html`, `/static/form.html`, etc.

- **Arithmetic APIs**
  - (Assuming REST endpoints are implemented in `main.go` for multiplication/division; see code for exact routes and payloads.)

### Example: Using the Linked List

```go
import "Go-First-project/datastructures"

ll := datastructures.NewLinkedList()
ll.Add(10)
ll.AddAtBeginning(5)
ll.Traverse() // Output: LinkedList: 5 -> 10
ll.Delete(10)
ll.Traverse() // Output: LinkedList: 5
```

### Example: Multiplication

```go
result := BasicMultiply(2.5, 4.0)
// result.Result == 10.0, result.Overflow == false
```

### Example: Division

```go
res, err := BasicDivide(10, 2)
// res.Result == 5.0, err == nil
```

---

## Testing

Run all tests with:

```sh
go test ./...
```

This will execute tests for handlers and arithmetic logic.

---

## License

This project is for educational purposes. See repository for license details if provided.

---

## Contributing

Pull requests and suggestions are welcome! Please open an issue or PR for improvements.

---
