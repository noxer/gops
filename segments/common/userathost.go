package common

import (
	"os"
	"os/user"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddUserAtHost adds the current user and the hostname of the current machine to the prompt. The compact flag trims all but the first segment of the host name.
func AddUserAtHost(segs []separator.Segment, compact bool) []separator.Segment {

	// get the hose name
	h, err := os.Hostname()
	if err != nil {
		// did not work? just place a ?
		h = "?"
	}

	if compact {
		hosts := strings.SplitN(h, ".", 2)
		if len(hosts) > 0 {
			h = hosts[0]
		}
	}

	// get the current user info
	u, err := user.Current()
	if err != nil {
		// if it didn't work place a ?
		u.Username = "?"
	}

	// create and add the host name segment
	s := separator.Segment{
		Text:       " " + u.Username + "@" + h + " ",
		Foreground: color.Black,
		Background: color.Green,
	}
	// if the current user is root, turn the segment red
	if u.Uid == "0" {
		s.Foreground = color.White
		s.Background = color.Red
	}

	return append(segs, s)

}
