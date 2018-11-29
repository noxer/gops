package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
	"gopkg.in/src-d/go-git.v4"
)

func main() {

	u, _ := user.Current()
	h, _ := os.Hostname()

	hostSeg := separator.Segment{
		Text:       " " + h + " ",
		Foreground: color.LightGray,
		Background: color.DarkGray,
	}

	userSeg := separator.Segment{
		Text:       " " + u.Username + " ",
		Foreground: color.Black,
		Background: color.LightGreen,
	}
	if u.Uid == "0" {
		userSeg.Foreground = color.White
		userSeg.Background = color.Red
	}

	endSeg := separator.Segment{
		Text:       " ",
		Foreground: color.Default,
		Background: color.Default,
	}

	segs := []separator.Segment{hostSeg, userSeg}
	segs = dirSegments(segs)
	segs = gitSegments(segs)
	segs = append(segs, endSeg)

	fmt.Print(separator.Render(segs...))

}

func dirSegments(segs []separator.Segment) []separator.Segment {

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

func gitSegments(segs []separator.Segment) []separator.Segment {

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
		return segs
	}

	segs = append(segs, separator.Segment{
		Text:       " \ue0a0 " + head.Name().Short() + " ",
		Foreground: color.DarkGray,
		Background: color.Yellow,
	})

	return segs

}
