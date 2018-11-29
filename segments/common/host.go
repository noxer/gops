package common

import (
	"os"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

func AddHost(segs []separator.Segment) []separator.Segment {

	h, err := os.Hostname()
	if err != nil {
		h = "?"
	}

	s := separator.Segment{
		Text:       " " + h + " ",
		Foreground: color.LightGray,
		Background: color.DarkGray,
	}

	return append(segs, s)

}
