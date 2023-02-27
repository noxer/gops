package common

import (
	"os"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddHost adds the host name of the current machine to the prompt. The compact flag trims all but the first segment of the host name.
func AddHost(segs []separator.Segment, compact bool) []separator.Segment {

	// get the host name
	h, err := os.Hostname()
	if err != nil {
		// did not work? just place a questionmark
		h = "?"
	}

	if compact {
		parts := strings.SplitN(h, ".", 2)
		if len(parts) > 0 {
			h = parts[0]
		}
	}

	// create and add the host name segment
	s := separator.Segment{
		Text:       " " + h + " ",
		Foreground: color.LightGray,
		Background: color.Black,
	}

	return append(segs, s)

}
