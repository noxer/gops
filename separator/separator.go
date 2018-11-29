package separator

import (
	"strings"

	"github.com/noxer/gops/color"
)

const (
	branch       = '\ue0a0'
	lock         = '\ue0a2'
	sepRight     = '\ue0b0'
	sepThinRight = '\ue0b1'
	sepLeft      = '\ue0b2'
	sepThinLeft  = '\ue0b3'
)

type Segment struct {
	Text       string
	Foreground color.Color
	Background color.Color
}

var EndSegment = Segment{
	Foreground: color.Default,
	Background: color.Default,
}

func Render(segs ...Segment) string {

	if len(segs) == 0 {
		return ""
	}

	buf := &strings.Builder{}

	buf.WriteString(color.Render(segs[0].Foreground, segs[0].Background))
	buf.WriteString(segs[0].Text)

	for i, s := range segs[1:] {

		if segs[i].Background == s.Background {
			// Both backgrounds are the same, render just a thin sep
			buf.WriteRune(sepThinRight)
		} else {
			buf.WriteString(color.Render(segs[i].Background, s.Background))
			buf.WriteRune(sepRight)
		}

		buf.WriteString(color.Render(s.Foreground, s.Background))
		buf.WriteString(s.Text)

	}

	return buf.String()

}
