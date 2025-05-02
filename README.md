# mars-go-tests

![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Tests](https://github.com/slazarska/mars-go-tests/actions/workflows/go-tests.yml/badge.svg)

A small educational Go project for testing Mars Rover Photos API by NASA.

## 🛰️ Project Description

This project demonstrates how to build and test a simple HTTP client in Go. 
It interacts with the [NASA Mars Rover Photos API](https://api.nasa.gov/) to fetch images taken by the rovers *Curiosity*, *Opportunity*, and *Spirit*.

It includes:

- Real API integration
- Struct-based JSON decoding
- Table-driven integration tests
- Mock-tests
- Logging with slog
- Allure reports
- GitHub Actions workflow for CI (automated testing)

## 📁 Project Structure
```
mars-go-tests/
├── cmd/                 # Entry point (optional CLI usage)
├── internal/
│   ├── api/             # API client logic
│   ├── config/          # Configuration loading
│   ├── constants/       # Base URL and other constants
│   ├── models/          # Structs matching NASA's JSON format
├── test/                # Integration tests using real API
├── go.mod / go.sum
└── README.md
```

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/mars-go-tests.git
cd mars-go-tests
```

### 2. Get your NASA API key
Register at https://api.nasa.gov

You'll get a free API key by email

### 3. Add your config
Create a file at internal/config/config.json with your API key:

```
{
  "api_key": "your_actual_nasa_api_key_here"
}
```
Important: This file is ignored by Git. Do not commit your API key.

### 4. Run the tests
```bash
go test ./...
```
You'll see real-time responses from NASA's API being tested.

### Example Test
Here's what one test looks like:
```
result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000")
assert.NoError(t, err)
assert.Greater(t, len(result.Photos), 0)
```

### Dependencies
- [Testify](https://github.com/stretchr/testify) for assertions
- Standard Go modules (`go.mod`, `go.sum`)

### License
MIT — feel free to use for educational purposes.

### Author
Maintained by [@slazarska](https://github.com/slazarska)