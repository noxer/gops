package git

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
	headName := filepath.Join(repoDir, "HEAD")

	p, err := ioutil.ReadFile(headName)
	if err != nil {
		return "?"
	}

	head := string(bytes.TrimSpace(p))

	if strings.HasPrefix(head, "ref: refs/heads/") {
		return head[16:]
	}

	// this seems to be just a hash. Try to find it in the packed refs
	pr := parsePackedRefs(repoDir)
	if pr == nil {
		return head
	}

	tag, ok := pr[head]
	if !ok {
		return head
	}

	return strings.TrimPrefix(tag, "refs/tags/")
}

func parsePackedRefs(repoDir string) map[string]string {
	prName := filepath.Join(repoDir, "packed-refs")
	f, err := os.Open(prName)
	if err != nil {
		return nil
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	m := make(map[string]string)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		// strings.Cut is part of Go 1.18 and not yet released
		// hash, ref, ok := strings.Cut(line, " ")
		// if !ok {
		// 	continue
		// }
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}
		hash, ref := parts[0], parts[1]

		m[hash] = ref
	}

	return m
}
