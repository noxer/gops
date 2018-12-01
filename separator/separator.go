package separator

import (
	"strings"

	"github.com/noxer/gops/color"
	"github.com/noxer/gops/symbols"
)

// Segment contains the information required to render a segment for the prompt
type Segment struct {
	Text       string
	Foreground color.Color
	Background color.Color
	Bold       bool
}

// EndSegment resets the color and formatting to default at the end of the prompt
var EndSegment = Segment{
	Foreground: color.Default,
	Background: color.Default,
}

// Render takes a list of segments and renders them into a prompt string
func Render(segs ...Segment) string {

	if len(segs) == 0 {
		return ""
	}

	// use a string builder to speed up the process
	buf := &strings.Builder{}

	// the first segment is just rendered as-is with no starting separator
	buf.WriteString(color.Render(segs[0].Foreground, segs[0].Background))
	buf.WriteString(segs[0].Text)

	for i, s := range segs[1:] {

		// render the separator
		if segs[i].Background == s.Background {
			// if both backgrounds are the same color, render just a thin separator
			buf.WriteRune(symbols.SepThinRight)
		} else {
			buf.WriteString(color.Render(segs[i].Background, s.Background))
			buf.WriteRune(symbols.SepRight)
		}

		// add the formatting to be used, written like this with extensibility in mind
		var fmts []color.Formatting
		if s.Bold {
			fmts = append(fmts, color.Bold)
		}

		// render the segment text
		buf.WriteString(color.Render(s.Foreground, s.Background, fmts...))
		buf.WriteString(s.Text)

	}

	// turn the buffer into a string and return it
	return buf.String()

}
