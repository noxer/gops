package example

import (
	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// Add an example segment to the segment list
func Add(segs []separator.Segment) []separator.Segment {

	text := "example"

	s := separator.Segment{
		Text:       " " + text + " ",
		Foreground: color.Black,
		Background: color.White,
	}

	return append(segs, s)

}
