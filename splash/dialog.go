package splash

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Convert hex color string to ANSI escape code (defined in splash.go)
// func hexToColor(hex string) (string, error) { ... } // This function is now in splash.go

// Dialog simulates a dialog in the terminal.
func (ui *Splash) Dialog(txt string, user bool, xdg string) (r bool) {
	if !ui.Config.Enabled {
		fmt.Printf("Dialog: %s\n", txt)
		return
	}
	fgColor, _ := hexToColor(ui.Config.FgColor)

	clearScreen := "\033[H\033[2J"
	fmt.Print(clearScreen)

	// Print dialog text
	fmt.Printf("%s%s\033[0m\n", fgColor, txt)

	// Print buttons
	if !user {
		if xdg == "" {
			fmt.Println("Press [Enter] to continue...")
		} else {
			fmt.Printf("Press [1] to view more info or [2] to continue...\n")
		}
	} else {
		fmt.Printf("Press [1] Yes or [2] No")
		if xdg != "" {
			fmt.Printf(" or [3] More Info")
		}
		fmt.Println()
	}

	// Handle user input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch {
	case input == "1" && user:
		r = true
		ui.Close()
	case input == "2" && user:
		ui.Close()
	case input == "1" && !user && xdg != "":
		fmt.Printf("Opening more info at: %s\n", xdg)
		// Implement xdg-open logic if needed
	case input == "2" && !user:
		ui.Close()
	case input == "3" && !user && xdg != "":
		fmt.Printf("Opening more info at: %s\n", xdg)
		// Implement xdg-open logic if needed
	default:
		fmt.Println("Invalid option")
	}
	return
}
