package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/segments/common"
	"github.com/noxer/gops/segments/git"
	"github.com/noxer/gops/separator"
	"github.com/noxer/gops/symbols"
)

var end = separator.Segment{
	Text:       " ",
	Foreground: color.Default,
	Background: color.Default,
}

func main() {

	prompt := flag.String("p", "PS1", "prompt type to render (PS1 or PS2)")
	shell := flag.String("s", "bash", "shell the prompt will be rendered for (bash or zsh)")
	flag.Parse()

	switch strings.ToUpper(*prompt) {
	case "PS1", "1":
		ps1(*shell)

	case "PS2", "2":
		ps2(*shell)

	default:
		flag.Usage()
		os.Exit(1)
	}

}

// ps1 renders the normal prompt
func ps1(shell string) {

	escape := initializeEscape(shell)

	var segs []separator.Segment
	segs = common.AddUserAtHost(segs, true)
	segs = common.AddDir(segs, true)
	segs = git.Add(segs)
	segs = append(segs, end)

	fmt.Print(separator.Render(escape, segs...))

}

// ps2 renders the continuation prompt
func ps2(shell string) {

	escape := initializeEscape(shell)

	segs := []separator.Segment{
		{Text: " " + string(symbols.Ellipsis) + " ", Foreground: color.LightGray, Background: color.DarkGray},
		end,
	}

	fmt.Println(separator.Render(escape, segs...))

}

func initializeEscape(shell string) func(string) string {
	switch shell {
	case "bash":
		return color.EscapeBASH
	case "zsh":
		return color.EscapeZSH
	}

	// no matching shell found, return empty escape function
	return func(s string) string {
		return s
	}
}
