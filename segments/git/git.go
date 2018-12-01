package git

import (
	"os"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
	"github.com/noxer/gops/symbols"
	"gopkg.in/src-d/go-git.v4"
)

// Add the Git branch as a segment
func Add(segs []separator.Segment) []separator.Segment {

	// get the current directory
	wd, err := os.Getwd()
	if err != nil {
		return segs
	}

	// search for a Git repo to we may be in
	repo, err := git.PlainOpenWithOptions(wd, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return segs
	}

	// use the HEAD of the repo to find the current branch
	head, err := repo.Head()
	if err != nil {
		// no HEAD in this repo, must be empty
		return append(segs, separator.Segment{
			Text:       " " + string(symbols.Branch) + " ? ",
			Foreground: color.DarkGray,
			Background: color.Yellow,
			Bold:       true,
		})
	}

	// add the segment with the repo name
	s := separator.Segment{
		Text:       " " + string(symbols.Branch) + " " + head.Name().Short() + " ",
		Foreground: color.DarkGray,
		Background: color.Yellow,
		Bold:       true,
	}

	return append(segs, s)

}
