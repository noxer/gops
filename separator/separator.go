package separator

import (
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/symbols"
)

type Segment struct {
	Text       string
	Foreground color.Color
	Background color.Color
	Bold       bool
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
			buf.WriteRune(symbols.SepThinRight)
		} else {
			buf.WriteString(color.Render(segs[i].Background, s.Background))
			buf.WriteRune(symbols.SepRight)
		}

		var fmts []color.Formatting
		if s.Bold {
			fmts = append(fmts, color.Bold)
		}

		buf.WriteString(color.Render(s.Foreground, s.Background, fmts...))
		buf.WriteString(s.Text)

	}

	return buf.String()

}
