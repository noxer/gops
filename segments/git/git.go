package git

import (
	"os"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
	"gopkg.in/src-d/go-git.v4"
)

func Add(segs []separator.Segment) []separator.Segment {

	wd, err := os.Getwd()
	if err != nil {
		return segs
	}

	repo, err := git.PlainOpenWithOptions(wd, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return segs
	}

	head, err := repo.Head()
	if err != nil {
		return append(segs, separator.Segment{
			Text:       " \ue0a0 ? ",
			Foreground: color.DarkGray,
			Background: color.Yellow,
		})
	}

	s := separator.Segment{
		Text:       " \ue0a0 " + head.Name().Short() + " ",
		Foreground: color.DarkGray,
		Background: color.Yellow,
	}

	return append(segs, s)

}
