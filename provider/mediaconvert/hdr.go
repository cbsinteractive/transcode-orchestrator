package mediaconvert

import (
	"fmt"
	"strconv"
	"strings"
)

type masterDisplay struct {
	// R, G, B color primaries (x,y) usually ordered as GBR in the string
	G, B, R,
	WP, // white point (x,y)
	L pt // luminance (max, min)
}

func (m masterDisplay) String() string {
	return fmt.Sprintf(
		"G%sB%sR%sWP%sL%s",
		m.G, m.B, m.R, m.WP, m.L,
	)
}

type pt struct {
	x, y int64
}

func (p pt) min() int64     { return p.y } // y is min
func (p pt) max() int64     { return p.x } // x is max
func (p pt) String() string { return fmt.Sprintf("(%d,%d)", p.x, p.y) }

var parseMasterDisplay = parseMasterDisplayFast

func parseMasterDisplayFast(s string) (d masterDisplay, err error) {
	const (
		tuples = 5     // G B R WP L
		delims = "()," // %s(%d,%d) is split across '(' and ',' and ')'
	)
	for _, r := range delims {
		if strings.Count(s, string(r)) != tuples {
			return d, fmt.Errorf("bad display string: %q", s)
		}
	}
	a := strings.FieldsFunc(s, func(r rune) bool { return strings.ContainsRune(delims, r) })
	if len(a) != tuples*len(delims) {
		return d, fmt.Errorf("too short: %d", len(a))
	}
	m := map[string]pt{}
	for i := 0; i < len(a); i += len(delims) {
		p := pt{}
		if p.x, err = strconv.ParseInt(a[i+1], 10, 64); err != nil {
			return
		}
		if p.y, err = strconv.ParseInt(a[i+2], 10, 64); err != nil {
			return
		}
		m[a[i]] = p
	}
	return masterDisplay{
		R:  pt{m["R"].x, m["R"].y},
		G:  pt{m["G"].x, m["G"].y},
		B:  pt{m["B"].x, m["B"].y},
		WP: pt{m["WP"].x, m["WP"].y},
		L:  pt{m["L"].x, m["L"].y},
	}, nil
}
