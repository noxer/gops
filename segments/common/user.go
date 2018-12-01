package common

import (
	"os/user"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

// AddUser adds the current user name segment
func AddUser(segs []separator.Segment) []separator.Segment {

	// get the current user info
	u, err := user.Current()
	if err != nil {
		return segs
	}

	// create the segment
	s := separator.Segment{
		Text:       " " + u.Username + " ",
		Foreground: color.Black,
		Background: color.LightGreen,
	}
	// if the current user is root, turn the segment red
	if u.Uid == "0" {
		s.Foreground = color.White
		s.Background = color.Red
	}

	return append(segs, s)

}
