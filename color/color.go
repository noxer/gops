package color

import "fmt"

// Color represents a terminal color
type Color byte

const (
	// Black color constant
	Black Color = 30
	// Red color constant
	Red Color = 31
	// Green color constant
	Green Color = 32
	// Yellow color constant
	Yellow Color = 33
	// Blue color constant
	Blue Color = 34
	// Magenta color constant
	Magenta Color = 35
	// Cyan color constant
	Cyan Color = 36
	// LightGray color constant
	LightGray Color = 37
	// Default color constant (default of the terminal)
	Default Color = 39
	// DarkGray color constant
	DarkGray Color = 90
	// LightRed color constant
	LightRed Color = 91
	// LightGreen color constant
	LightGreen Color = 92
	// LightYellow color constant
	LightYellow Color = 93
	// LightBlue color constant
	LightBlue Color = 94
	// LightMagenta color constant
	LightMagenta Color = 95
	// LightCyan color constant
	LightCyan Color = 96
	// White color constant
	White Color = 97
)

// Formatting represents a character formatting
type Formatting byte

const (
	// Reset formatting constant
	Reset Formatting = 0
	// Bold formatting constant
	Bold Formatting = 1
	// Dim formatting constant
	Dim Formatting = 2
	// Underlined formatting constant
	Underlined Formatting = 4
	// Blink formatting constant
	Blink Formatting = 5
	// Reverse (inverted colors) formatting constant
	Reverse Formatting = 7
	// Hidden formatting constant
	Hidden Formatting = 8
	// ResetBold formatting constant
	ResetBold Formatting = 21
	// ResetDim formatting constant
	ResetDim Formatting = 22
	// ResetUnderlined formatting constant
	ResetUnderlined Formatting = 24
	// ResetBlink formatting constant
	ResetBlink Formatting = 25
	// ResetReverse formatting constant
	ResetReverse Formatting = 27
	// ResetHidden formatting constant
	ResetHidden Formatting = 28
)

func EscapeBASH(s string) string {
	return "\001" + s + "\002"
}

func EscapeZSH(s string) string {
	return "%{" + s + "%}"
}

// Render takes the foreground and background color as well as additional formatting constants and generates a terminal command.
func Render(fg, bg Color, escape func(string) string, fs ...Formatting) string {
	buf := ""
	for _, f := range fs {
		buf += fmt.Sprintf(";%d", f)
	}

	return escape(fmt.Sprintf("\033[0;%d;%d%sm", fg, bg+10, buf))
}
