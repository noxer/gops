package common

import (
	"os/user"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

func AddUser(segs []separator.Segment) []separator.Segment {

	u, err := user.Current()
	if err != nil {
		return segs
	}

	s := separator.Segment{
		Text:       " " + u.Username + " ",
		Foreground: color.Black,
		Background: color.LightGreen,
	}
	if u.Uid == "0" {
		s.Foreground = color.White
		s.Background = color.Red
	}

	return append(segs, s)

}
