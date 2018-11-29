package main

import (
	"fmt"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/segments/common"
	"github.com/noxer/gops/segments/git"
	"github.com/noxer/gops/separator"
)

func main() {

	end := separator.Segment{
		Text:       " ",
		Foreground: color.Default,
		Background: color.Default,
	}

	var segs []separator.Segment
	segs = common.AddHost(segs)
	segs = common.AddUser(segs)
	segs = common.AddDir(segs)
	segs = git.Add(segs)
	segs = append(segs, end)

	fmt.Print(separator.Render(segs...))

}
