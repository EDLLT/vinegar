package splash

import (
	"fmt"
	"strings"
)

type Style int

const (
	Compact Style = iota
	Familiar
)

func (s Style) Size() (w, h int) {
	switch s {
	case Compact:
		w = 50
		h = 10
	case Familiar:
		w = 60
		h = 20
	}
	return
}

func (ui *Splash) drawCompact() {
	width, _ := ui.Style.Size() // Removed unused height
	clearScreen := "\033[H\033[2J"
	fmt.Print(clearScreen)

	fmt.Println(strings.Repeat("=", width))
	fmt.Printf("Logo: %s\n", ui.drawLogo())
	fmt.Printf("Message: %s\n", ui.Message)
	fmt.Printf("Description: %s\n", ui.drawDesc())
	progressBar := ui.drawProgressBar()
	fmt.Printf("Progress: %s\n", progressBar)
	fmt.Println(strings.Repeat("=", width))
	ui.drawButtons()
}

func (ui *Splash) drawFamiliar() {
	width, _ := ui.Style.Size() // Removed unused height
	clearScreen := "\033[H\033[2J"
	fmt.Print(clearScreen)

	fmt.Println(strings.Repeat("=", width))
	fmt.Printf("Logo: %s\n", ui.drawLogo())
	fmt.Printf("Message: %s\n", ui.Message)
	fmt.Printf("Description: %s\n", ui.drawDesc())
	progressBar := ui.drawProgressBar()
	fmt.Printf("Progress: %s\n", progressBar)
	fmt.Println(strings.Repeat("=", width))
	ui.drawButtons()
}

func (ui *Splash) drawProgressBar() string {
	total := 20
	completed := int(ui.Progress * float32(total))
	bar := fmt.Sprintf("[%s%s]", strings.Repeat("#", completed), strings.Repeat("-", total-completed))
	return bar
}

// These functions are now in widgets.go:
// func (ui *Splash) drawLogo() string {
// 	return "LOGO"
// }
// 
// func (ui *Splash) drawDesc() string {
// 	return "Description here"
// }
