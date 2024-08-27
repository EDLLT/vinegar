package splash

import (
	"gioui.org/layout"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// The declaration of selectedButton is now in splash.go
// var selectedButton int

// func (ui *Splash) drawLogo() string {
// 	return ui.Logo
// }

// func (ui *Splash) drawButtons() {
// 	// Styling for buttons (optional, using ANSI escape codes)
// 	bgColor, _ := hexToColor("#000000") // Black background
// 	fgColor, _ := hexToColor("#FFFFFF") // White foreground

// 	if ui.LogPath != "" {
// 		fmt.Printf("%s[%s]%s\n", bgColor, "1. Show logs", fgColor)
// 	}
// 	fmt.Printf("%s[%s]%s\n", bgColor, "2. Cancel", fgColor)

// 	// // Read user input
// 	// reader := bufio.NewReader(os.Stdin)
// 	// input, _ := reader.ReadString('\n')
// 	// input = strings.TrimSpace(input)

// 	// // Determine the selected button based on input
// 	// switch input {
// 	// case "1":
// 	// 	selectedButton = 1 // Show logs
// 	// case "2":
// 	// 	selectedButton = 2 // Cancel
// 	// default:
// 	// 	fmt.Println("Invalid option.")
// 	// }
// }

// func (ui *Splash) drawDesc() string {
// 	return ui.Desc
// }
