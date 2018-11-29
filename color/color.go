package color

import "fmt"

type Color byte

const (
	Black        Color = 30
	Red          Color = 31
	Green        Color = 32
	Yellow       Color = 33
	Blue         Color = 34
	Magenta      Color = 35
	Cyan         Color = 36
	LightGray    Color = 37
	Default      Color = 39
	DarkGray     Color = 90
	LightRed     Color = 91
	LightGreen   Color = 92
	LightYellow  Color = 93
	LightBlue    Color = 94
	LightMagenta Color = 95
	LightCyan    Color = 96
	White        Color = 97
)

type Formatting byte

const (
	Reset           Formatting = 0
	Bold            Formatting = 1
	Dim             Formatting = 2
	Underlined      Formatting = 4
	Blink           Formatting = 5
	Reverse         Formatting = 7
	Hidden          Formatting = 8
	ResetBold       Formatting = 21
	ResetDim        Formatting = 22
	ResetUnderlined Formatting = 24
	ResetBlink      Formatting = 25
	ResetReverse    Formatting = 27
	ResetHidden     Formatting = 28
)

func Render(fg, bg Color, fs ...Formatting) string {
	buf := ""
	for _, f := range fs {
		buf += fmt.Sprintf(";%d", f)
	}

	return fmt.Sprintf("\001\033[%d;%d%sm\002", fg, bg+10, buf)
}
