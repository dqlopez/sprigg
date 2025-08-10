Project: Sprigg - Chrono Cross Search Automator
Overview
The goal of this project is to create a command-line utility in Go that automates a simple web search. The script will randomly select a character from the game Chrono Cross and perform a search for that character on the Bing search engine using the Microsoft Edge browser.

Requirements
The script must be written entirely in the Go programming language.

It must be a command-line application that can be executed directly from a terminal.

The application must open the Microsoft Edge browser.

The search query must be for a randomly selected Chrono Cross character.

The search engine used must be Bing.

Technical Details
Language: Go

Dependencies: No external libraries are strictly required for the core functionality. The standard library's os/exec package can be used to execute system commands to open the browser.

Cross-platform Considerations: The script's os/exec command will need to be written to be compatible with different operating systems (Windows, macOS, and Linux) to reliably open Microsoft Edge. A simple if-else block can handle this by checking the operating system.

Core Logic
The script should follow these steps:

Define an array of strings containing the names of various Chrono Cross characters (e.g., "Serge", "Kid", "Sprigg", "Glenn", "Harle", etc.).

Generate a random number to select one of the character names from the array.

Construct the full search URL, for example: https://www.bing.com/search?q=chrono+cross+{character_name}. The character name should be URL-encoded for proper formatting.

Use os/exec to run the appropriate command for the detected operating system to open the constructed URL in Microsoft Edge. For example:

Windows: Use the start command with the microsoft-edge: URI scheme.

macOS: Use the open command with the -a flag to specify "Microsoft Edge".

Linux: Use the xdg-open command, which typically respects the user's default browser settings but can be configured.