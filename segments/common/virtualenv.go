package common

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddVirtualenv adds the name of the current virtualenv to the prompt
func AddVenv(segs []separator.Segment) []separator.Segment {

	// get the name of the current virtualenv
	venvpath := os.Getenv("VIRTUAL_ENV")
	if venvpath == "" {
		return segs

	}

	// strip the full path
	parts := strings.Split(venvpath, string(filepath.Separator))
	venv := parts[len(parts)-1]

	// create the segment
	s := separator.Segment{
		Text:       " (" + venv + ") ",
		Foreground: color.White,
		Background: color.DarkGray,
	}

	return append(segs, s)

}
