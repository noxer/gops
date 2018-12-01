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

	prompt := ""
	flag.StringVar(&prompt, "p", "PS1", "prompt type to render (PS1 or PS2)")
	flag.Parse()

	switch strings.ToUpper(prompt) {
	case "PS1", "1":
		ps1()

	case "PS2", "2":
		ps2()

	default:
		flag.Usage()
		os.Exit(1)
	}

}

// ps1 renders the normal prompt
func ps1() {

	var segs []separator.Segment
	segs = common.AddHost(segs)
	segs = common.AddUser(segs)
	segs = common.AddDir(segs)
	segs = git.Add(segs)
	segs = append(segs, end)

	fmt.Print(separator.Render(segs...))

}

// ps2 renders the continuation prompt
func ps2() {

	segs := []separator.Segment{
		{Text: " " + string(symbols.Ellipsis) + " ", Foreground: color.LightGray, Background: color.DarkGray},
		end,
	}

	fmt.Print(separator.Render(segs...))

}
