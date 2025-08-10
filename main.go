package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"strings"
	"runtime"
)

// List of Chrono Cross characters
var characters = []string{
	"Serge", "Kid", "Sprigg", "Lynx", "Harle",
	"Glenn", "Leena", "Guile", "Karsh", "Riddel",
	"Zoah", "Marcy", "Viper", "Razzly", "Poshul",
	"Fargo", "Orlha", "Irenes", "Pierre", "Steena",
	"Van", "Janice", "Macha", "Korcha", "Mel",
	"Greco", "Skelly", "Funguy", "Leah", "Orcha",
	"Draggy", "Starky", "Razzly", "Radius", "Miki",
	"Zappa", "Lucca", "Norris", "Grobyc",
}

// List of game-related terms to add variety to searches
var gameTerms = []string{
	"element", "tech", "summon", "sword", "dragon", "island",
	"village", "temple", "dungeon", "boss", "weapon", "armor",
	"accessory", "item", "magic", "spell", "ability", "skill",
	"quest", "story", "ending", "character", "party", "ally",
	"enemy", "villain", "hero", "heroine", "world", "map",
	"puzzle", "battle", "combat", "strategy", "guide", "walkthrough",
	"artwork", "music", "soundtrack", "remaster", "remake", "sequel",
	"easter egg", "secret", "hidden", "unlockable", "achievement", "trophy",
}

func checkMacOSPermissions() {
	// Check if Terminal has accessibility permissions
	cmd := exec.Command("osascript", "-e", `
		tell application "System Events"
			tell process "Terminal"
				try
					set frontmost to true
					true
				on error
					return false
				end try
			end tell
		end tell
	`)

	if err := cmd.Run(); err != nil {
		showPermissionInstructions()
		os.Exit(1)
	}
}

func showPermissionInstructions() {
	fmt.Println("")
	fmt.Println("⚠️  Permission Required ⚠️")
	fmt.Println("This app needs accessibility permissions to control Microsoft Edge.")
	fmt.Println("Please follow these steps to enable it:")
	fmt.Println("1. Open System Preferences")
	fmt.Println("2. Go to Security & Privacy")
	fmt.Println("3. Select the Privacy tab")
	fmt.Println("4. Select Accessibility from the left sidebar")
	fmt.Println("5. Click the lock icon and enter your password")
	fmt.Println("6. Click the + button and add Terminal (or your IDE) to the list")
	fmt.Println("7. Make sure the checkbox is checked")
	fmt.Println("8. Run this program again")
	fmt.Println("")
}

func main() {
	// Parse command line flags
	count := flag.Int("n", 1, "Number of searches to perform (default: 1)")
	flag.Parse()

	// Validate count
	if *count < 1 {
		fmt.Println("Number of searches must be at least 1")
		os.Exit(1)
	}

	// Check for required permissions on macOS
	if runtime.GOOS == "darwin" {
		checkMacOSPermissions()
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Open Microsoft Edge to bing.com
	openEdge()

	// Perform the specified number of searches
	for i := 0; i < *count; i++ {
		// Select a random character and game term
		character := characters[rand.Intn(len(characters))]
		term := gameTerms[rand.Intn(len(gameTerms))]
		searchQuery := fmt.Sprintf("Chrono Cross %s %s", character, term)

		fmt.Printf("Searching for: %s (%d/%d)\n", character, i+1, *count)

		// Use AppleScript to simulate keyboard input
		applescript := fmt.Sprintf(`
		tell application "Microsoft Edge"
			activate
		end tell

		delay 1 -- Wait for Edge to activate (reduced from 2s)

		tell application "System Events"
			-- Make sure we're on the frontmost application
			set frontApp to name of first application process whose frontmost is true
			if frontApp is not "Microsoft Edge" then
				tell application "Microsoft Edge" to activate
				delay 1
			end if

			-- Focus the address bar using Cmd+L
			key code 37 using command down
			delay 0.25 -- Reduced from 0.5s

			-- Type the search query directly in the address bar
			keystroke "%s"
			delay 0.25 -- Reduced from 0.5s

			-- Press Enter to search
			key code 36

			delay 2 -- Wait for search results
		end tell
		`, strings.ReplaceAll(searchQuery, "'", "'\\''"))

		// Execute the AppleScript
		cmd := exec.Command("osascript", "-e", applescript)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing AppleScript: %v\nOutput: %s\n", err, output)
			if strings.Contains(string(output), "not allowed to send keystrokes") {
				showPermissionInstructions()
				os.Exit(1)
			}
		}

		// Wait between searches (3 seconds + random 1-2 seconds)
		if i < *count-1 {
			waitTime := 3 + time.Duration(rand.Intn(2))
			time.Sleep(waitTime * time.Second)
		}
	}
}

func openEdge() {
	// Simple function to open Edge with Bing
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "darwin":
		// On macOS, use open command
		cmd = exec.Command("open", "-a", "Microsoft Edge", "https://www.bing.com")
	case "windows":
		// On Windows, use start command
		cmd = exec.Command("cmd", "/c", "start", "msedge", "https://www.bing.com")
	default:
		// On Linux, try with microsoft-edge
		cmd = exec.Command("microsoft-edge", "https://www.bing.com")
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error opening Microsoft Edge: %v\n", err)
	}
	
	// Give the browser time to open (reduced from 3s)
	time.Sleep(1500 * time.Millisecond)
}
