package common

import (
	"os"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddHost adds the host name of the current machine to the prompt
func AddHost(segs []separator.Segment) []separator.Segment {

	// get the hose name
	h, err := os.Hostname()
	if err != nil {
		// did not work? just place a ?
		h = "?"
	} else {
		h = strings.TrimSuffix(h, ".local")
	}

	// create and add the host name segment
	s := separator.Segment{
		Text:       " " + h + " ",
		Foreground: color.LightGray,
		Background: color.DarkGray,
	}

	return append(segs, s)

}
