package splash

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrClosed = errors.New("window closed")

type Config struct {
	Enabled     bool   `toml:"enabled"`
	LogoPath    string `toml:"logo_path"`
	Style       string `toml:"style"`
	BgColor     string `toml:"background"`
	FgColor     string `toml:"foreground"`
	CancelColor string `toml:"cancel,red"`
	AccentColor string `toml:"accent"`
	TrackColor  string `toml:"track,gray1"`
	InfoColor   string `toml:"info,gray2"`
}

type Splash struct {
	Config   *Config
	Logo     string
	Style    Style  // Added Style field
	Message  string // Added Message field
	Desc     string
	Progress float32 // Added Progress field
	Closed   bool // Changed to Closed (capital C)
	LogPath  string // Added LogPath field
}

func hexToColor(hex string) (string, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return "", errors.New("invalid color format")
	}
	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return "", err
	}
	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b), nil
}

func New(cfg *Config) *Splash {
	if !cfg.Enabled {
		return &Splash{
			Closed: true,
			Config: cfg,
		}
	}

	logo := "[Logo Placeholder]"

	if cfg.LogoPath != "" {
		if data, err := os.ReadFile(cfg.LogoPath); err == nil {
			logo = string(data)
		}
	}

	return &Splash{
		Config:  cfg,
		Logo:    logo,
		Closed:  false,
	}
}

func (ui *Splash) SetMessage(msg string) {
	if ui.Closed {
		return
	}
	ui.Message = msg
	ui.render()
}

func (ui *Splash) SetDesc(desc string) {
	if ui.Closed {
		return
	}
	ui.Desc = desc
	ui.render()
}

func (ui *Splash) SetProgress(progress float32) {
	if ui.Closed {
		return
	}
	ui.Progress = progress
	ui.render()
}

func (ui *Splash) Close() {
	if ui.Closed {
		return
	}
	ui.Closed = true
	ui.render()
}

func (ui *Splash) IsClosed() bool {
	return ui.Closed
}

// Declare selectedButton at package level (to be accessible from both files)
var selectedButton int

func (ui *Splash) render() {
	if ui.Closed {
		return
	}

	fmt.Printf("\033[H\033[2J") // Clear screen

	bgColor, _ := hexToColor(ui.Config.BgColor)
	fgColor, _ := hexToColor(ui.Config.FgColor)

	// Print Logo
	fmt.Println(bgColor + ui.Logo + "\033[0m")

	// Print Message
	fmt.Println(fgColor + ui.Message + "\033[0m")

	// Print Description
	fmt.Println(fgColor + ui.Desc + "\033[0m")

	// Print Progress
	progressBar := int(ui.Progress * 50)
	fmt.Printf("Progress: [%s%s] %.0f%%\n", strings.Repeat("=", progressBar), strings.Repeat(" ", 50-progressBar), ui.Progress*100)

	// Call drawButtons function to handle button input
	ui.drawButtons()

	// Handle user input based on selectedButton
	switch selectedButton {
	case 1: // Show logs
		if ui.LogPath != "" {
			// Open the log file
		}
		selectedButton = 0 // Reset selectedButton
	case 2: // Cancel
		ui.Close()
		selectedButton = 0 // Reset selectedButton
		// ... (other cases)
	}
}

// Display the terminal-based splash screen (new function)
func (ui *Splash) Display() {
	// Clear the screen
	fmt.Printf("\033[H\033[2J")

	for !ui.IsClosed() {
		ui.render()
	}
}

// Run (kept for compatibility) - now calls Display
func (ui *Splash) Run() error {
	if ui.Closed { // Changed to Closed (capital C)
		return nil
	}

	ui.Display() // Call Display to handle the terminal-based splash

	return nil // Or return an error if needed
}

func main() {
	cfg := &Config{
		Enabled: true,
		LogoPath: "logo.txt",
		BgColor: "#000000", // Default black background
		FgColor: "#FFFFFF", // Default white foreground
	}
	splash := New(cfg)
	splash.SetMessage("Welcome to Vinegar")
	splash.SetDesc("Loading...")
	splash.SetProgress(0.5)

	// Start the splash screen
	splash.Run()
}
