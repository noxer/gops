package common

import (
	"os/exec"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddNodeVersion adds the version of the currently used node.js installation the prompt
func AddNodeVersion(segs []separator.Segment) []separator.Segment {

	out, err := exec.Command("node", "-v").Output()

	version := strings.TrimSpace(string(out))

	if err != nil {
		// did not work? just place a ?
		version = "?"
	}

	// create the segment
	s := separator.Segment{
		Text:       "node: " + version,
		Foreground: color.White,
		Background: color.DarkGray,
	}

	return append(segs, s)

}
