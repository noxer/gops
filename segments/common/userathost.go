package common

import (
	"os"
	"os/user"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddHost adds the host name of the current machine to the prompt
func AddUserAtHost(segs []separator.Segment) []separator.Segment {

	// get the hose name
	h, err := os.Hostname()
	if err != nil {
		// did not work? just place a ?
		h = "?"
	} else {
		h = strings.TrimSuffix(h, ".local")
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
		Foreground: color.LightGray,
		Background: color.Black,
	}
	// if the current user is root, turn the segment red
	if u.Uid == "0" {
		s.Foreground = color.White
		s.Background = color.Red
	}

	return append(segs, s)

}
