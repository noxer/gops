package common

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
	"github.com/noxer/gops/symbols"
)

func wd() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// if the home dir is a prefix of the current dir, replace it by ~
	user, err := user.Current()
	if err == nil {
		if strings.HasPrefix(wd, user.HomeDir) {
			wd = "~" + strings.TrimPrefix(wd, user.HomeDir)
		}
	}

	return wd, nil
}

// AddDir appends the directory segments for the current directory to the list of segments.
func AddDir(segs []separator.Segment, compact bool) []separator.Segment {

	wd, err := wd()
	if err != nil {
		return segs
	}

	if compact {
		s := separator.Segment{
			Text:       " " + wd + " ",
			Foreground: color.LightGray,
			Background: color.DarkGray,
			Bold:       true,
		}
		return append(segs, s)
	}

	// split the directory into segments
	parts := strings.Split(wd, string(filepath.Separator))

	// calculate the size of the resulting string
	i := len(parts) * 3
	for _, p := range parts {
		i += len(p)
	}

	// if the path started with a /, add it back (was removed when splitting)
	if parts[0] == "" {
		parts[0] = "/"
	}
	// if the path ended with a /, remove the last (empty) segment
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	// if the path is too long, maybe we can remove some of the segments in the middle
	if i > 20 && len(parts) > 5 {
		parts[2] = " " + string(symbols.Ellipsis) + " "
		copy(parts[3:], parts[len(parts)-2:])
		parts = parts[:5]
	}

	// add the path segments
	for _, p := range parts {
		s := separator.Segment{
			Text:       " " + p + " ",
			Foreground: color.LightGray,
			Background: color.DarkGray,
			Bold:       true,
		}
		segs = append(segs, s)
	}

	return segs

}
