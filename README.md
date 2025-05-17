# go-crawler
A simple recursive web crawler written in Go. Starts from a given URL (eventually), parses HTML content, and recursively visits all links up to a specified max depth

## Features
- Parses and extracts <a href=""> links
- skips fragment only URLs
- Prevents duplicate visits using a "visited" map
- Limits traversal depth

## How to use
1. Clone the repository:
   ```bash
   git clone https://github.com/scott-richardson-135/go-crawler.git
   cd go-crawler
   ```
2. Run the crawler:
   go run main.go

## Things I want to add
- Better code separation and internal structure
- Logging to file
- Concurrency with goroutines
- Command line arguments for url or max depth
