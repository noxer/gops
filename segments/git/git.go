package git

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/separator"
	"github.com/noxer/gops/symbols"
)

// Add the Git branch as a segment
func Add(segs []separator.Segment) []separator.Segment {

	branch := findGitRepo()
	if branch == "" {
		return segs
	}

	// add the segment with the repo name
	s := separator.Segment{
		Text:       " " + string(symbols.Branch) + " " + branch + " ",
		Foreground: color.DarkGray,
		Background: color.Yellow,
		Bold:       true,
	}

	return append(segs, s)

}

func findGitRepo() string {

	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {

		repoDir := filepath.Join(dir, ".git")
		if _, err = os.Stat(filepath.Join(dir, ".git")); err == nil {
			return readHEAD(repoDir)
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent

	}

	return ""

}

func readHEAD(repoDir string) string {

	repoDir = filepath.Join(repoDir, "HEAD")

	p, err := ioutil.ReadFile(repoDir)
	if err != nil {
		return "?"
	}

	return string(bytes.TrimPrefix(bytes.TrimSpace(p), []byte("ref: refs/heads/")))

}
