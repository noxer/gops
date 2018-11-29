package common

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
)

func AddDir(segs []separator.Segment) []separator.Segment {

	wd, err := os.Getwd()
	if err != nil {
		return segs
	}

	user, err := user.Current()
	if err == nil {
		if strings.HasPrefix(wd, user.HomeDir) {
			wd = "~" + strings.TrimPrefix(wd, user.HomeDir)
		}
	}

	parts := strings.Split(wd, string(filepath.Separator))

	if parts[0] == "" {
		parts[0] = "/"
	}
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	for _, p := range parts {
		s := separator.Segment{
			Text:       " " + p + " ",
			Foreground: color.LightGray,
			Background: color.DarkGray,
		}
		segs = append(segs, s)
	}

	return segs

}
