# Sprigg - Chrono Cross Search Automator

A simple command-line utility written in Go that searches for random Chrono Cross characters on Bing using Microsoft Edge.

## Features

- Randomly selects a Chrono Cross character
- Uses the system's default browser for searches
- Simple and reliable implementation
- Shows progress with the current search count (e.g., "1/5")
- Adds random delays between searches

## Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or later
- A web browser installed (any modern browser will work)

## Installation

1. Clone this repository
2. Build the application:
   ```bash
   go build -o sprigg
   ```

## Usage

Run the compiled binary:

```bash
# Basic usage (searches for 1 character)
./sprigg

# Search for multiple characters (e.g., 5)
./sprigg -n 5

# On Windows
sprigg.exe -n 3  # Searches for 3 characters
```

### Options:
- `-n` : Number of searches to perform (default: 1)
  
The application will:
1. Open your default browser with a Bing search for a random Chrono Cross character
2. Wait a few seconds between searches
3. Show progress with the current search count (e.g., "1/5")

## Building for Different Platforms

You can build the application for different operating systems using Go's cross-compilation:

```bash
# For Windows
GOOS=windows GOARCH=amd64 go build -o sprigg.exe

# For macOS
GOOS=darwin GOARCH=amd64 go build -o sprigg

# For Linux
GOOS=linux GOARCH=amd64 go build -o sprigg
```

## License

MIT
