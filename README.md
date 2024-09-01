# URL Shortener Service
This project is a simple URL shortener service built with Go. It provides an API to shorten URLs, retrieve the original URLs, and track metrics of the top shortened domains.

## Features
* __Shorten URL__: Accepts a URL and returns a shortened version using Base62 encoding.

* __Redirect__: Redirects the shortened URL to the original URL.

* __Metrics API__: Returns the top 3 most shortened domain names.

## Project Structure
The project is organized as follows:
```
url-shortener/
│
├── cmd/server/
│   └── main.go             # Entry point of the application
│
├── internal/
│   ├── api/                # Handlers, routes, and middlewares
│   ├── config/             # Loading configs like port number
│   ├── errors/             # Custom error logic
│   ├── models/             # Data structures and validation logic
│   ├── storage/            # Storage interface and in-memory implementation
│   ├── services/           # URL shortening and metrics logic
│   └── utils/              # Utility functions like Base62 encoding
│
└── go.mod                  # Go module file
```
## Prerequisites
Go 1.23 or higher
Docker (for containerization)

# Getting Started
## 1. Clone the repository
```
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```
## 2. Build and Run Locally
You can use the provided Makefile to manage the project.
### Format, Vendor, and Tidy Go Modules
```
make tidy
```
### Run the Application
```
make run
```
### Run Tests
```
make test
```
## 3. Build and Run with Docker
### Build the Docker Image
```
make docker-build
```
### Run the Docker Container
```
make docker-run
```
# API Usage
## Shorten a URL
```
curl --location 'localhost:8080/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www.google.com/search?q=rich+dad+poor+dad"
}'
```
```
{
    "shortened_url": "localhost:8080/r/AnXmfMMA"
}
```
## Redirect to Original URL
```
curl --location 'localhost:8080/r/AnXkH44A'
```
## Get Top 3 Shortened Domains
```
curl --location 'localhost:8080/metrics/top-domains'
```
```
{
    "top_domains": [
        {
            "DomainURL": "www.google.com",
            "VisitCount": 23
        },
        {
            "DomainURL": "www.facebook.com",
            "VisitCount": 10
        },
        {
            "DomainURL": "www.leetcode.com",
            "VisitCount": 8
        }
    ]
}
```
# Unique ID Generatation
Used twitter's snowflake approach to generate unique ID of 48 bits (41 bits for timestamp, 4 bits for instance ID and 3 bits for sequence ID which gets reset to 0 after end of millisecond).
Which is converted to base64 encoded string of 8 characters.
Total number of unique short URLs that can be generated with this.
```
64^8 = ~281 trillion
```
As timestamp is using 41 bits and epoch time is Jan 1, 2024. This approach is good for 69 years from the epoch time.
```
2^41 Millisecond
2^41/1000/60/60/24/365 = ~69.7 years
```
# Future Improvements
### 1. Database Integration:
Use a database like PostgreSQL or MySQL for persistent storage. This will allow the service to retain shortened URLs and metrics across restarts.
Include a docker-compose.yml file to easily set up the database along with the service.
### 2. Improved Logging:
Implement logging to a file for better traceability and debugging.
Use structured logging libraries like logrus or zap.
### 3. Authentication:
Add authentication (e.g., API keys, OAuth) for URL shortening to restrict usage.
### 4. Enhanced Comments:
Improve inline documentation and comments to make the codebase more understandable for new contributors.
### 5. Enchanced Unit Tests
Improve the test coverage and add more test cases for edge scenarios and error handling. 

# License
This project is licensed under the MIT License.
