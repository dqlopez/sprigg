# Sprigg - Chrono Cross Search Automator

A command-line utility written in Go that automates Bing searches for random Chrono Cross characters using Microsoft Edge. Perfect for earning Microsoft Rewards points.

## Features

- Randomly selects a Chrono Cross character
- Uses Microsoft Edge for searches (required for Microsoft Rewards)
- Simulates real user interaction with the search box
- Shows progress with the current search count (e.g., "1/5")
- Adds random delays between searches to mimic human behavior
- Cross-platform support (macOS, Windows, Linux)

## Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or later
- Microsoft Edge browser (required for Microsoft Rewards)

### macOS Permissions

On macOS, you'll need to grant accessibility permissions for the terminal/IDE to control Microsoft Edge:

1. Open **System Preferences** > **Security & Privacy** > **Privacy** tab
2. Select **Accessibility** from the left sidebar
3. Click the lock icon 🔒 and enter your password
4. Click **+** and add:
   - **Terminal** (if running from terminal)
   - Your **IDE** (if running from VS Code, etc.)
5. Make sure the checkbox is checked for the added application
6. Run the program again

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/sprigg.git
   cd sprigg
   ```

2. Build the application:
   ```bash
   go build -o sprigg
   ```
   
3. (Optional) Move to a directory in your PATH:
   ```bash
   sudo mv sprigg /usr/local/bin/
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
  
### How It Works

The application will:
1. Open Microsoft Edge and navigate to bing.com
2. Enter the search query in the address bar
3. Submit the search
4. Wait a few seconds between searches
5. Show progress with the current search count (e.g., "1/5")

### Important Notes

- Microsoft Edge must be installed
- On first run, you'll need to log in to your Microsoft account in Edge
- The script includes random delays to mimic human behavior

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
