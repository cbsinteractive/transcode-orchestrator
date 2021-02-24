package mediaconvert

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	masterDisplayValueGroupGreen      = "green"
	masterDisplayValueGroupBlue       = "blue"
	masterDisplayValueGroupRed        = "red"
	masterDisplayValueGroupWhitepoint = "whitepoint"
	masterDisplayValueGroupLuminance  = "luminance"
)

var (
	masterDisplayRegxp = regexp.MustCompile(`(?P<` + masterDisplayValueGroupGreen + `>G\(\d+,\d+\))|` +
		`(?P<` + masterDisplayValueGroupBlue + `>B\(\d+,\d+\))|` +
		`(?P<` + masterDisplayValueGroupRed + `>R\(\d+,\d+\))|` +
		`(?P<` + masterDisplayValueGroupWhitepoint + `>WP\(\d+,\d+\))|` +
		`(?P<` + masterDisplayValueGroupLuminance + `>L\(\d+,\d+\))`)

	nonNumericRegex = regexp.MustCompile(`[^0-9]+`)
)

type masterDisplay struct {
	redPrimaryX   int64
	redPrimaryY   int64
	greenPrimaryX int64
	greenPrimaryY int64
	bluePrimaryX  int64
	bluePrimaryY  int64
	whitePointX   int64
	whitePointY   int64
	maxLuminance  int64
	minLuminance  int64
}

type tuple struct {
	x, y int64
}

var parseMasterDisplay = parseMasterDisplayRegexp

// var parseMasterDisplay = parseMasterDisplayFast

func parseMasterDisplayRegexp(encoded string) (masterDisplay, error) {
	groupRegex := masterDisplayRegxp

	matchGroup := groupRegex.FindAllStringSubmatch(encoded, -1)

	results := map[string]tuple{}
	for _, matches := range matchGroup {
		for i, match := range matches {
			if match == "" {
				continue
			}

			coordinateParts := strings.Split(match, ",")
			if len(coordinateParts) != 2 {
				return masterDisplay{}, fmt.Errorf("invalid master display format: %q", encoded)
			}

			x, err := numbersInString(coordinateParts[0])
			if err != nil {
				return masterDisplay{}, errors.Wrap(err, "unable to parse x value to int")
			}

			y, err := numbersInString(coordinateParts[1])
			if err != nil {
				return masterDisplay{}, errors.Wrap(err, "unable to parse y value to int")
			}

			results[groupRegex.SubexpNames()[i]] = tuple{x: x, y: y}
		}
	}
	if len(results) != 6 {
		return masterDisplay{}, fmt.Errorf("invalid master display format: %q", encoded)
	}

	return masterDisplay{
		redPrimaryX:   results[masterDisplayValueGroupRed].x,
		redPrimaryY:   results[masterDisplayValueGroupRed].y,
		greenPrimaryX: results[masterDisplayValueGroupGreen].x,
		greenPrimaryY: results[masterDisplayValueGroupGreen].y,
		bluePrimaryX:  results[masterDisplayValueGroupBlue].x,
		bluePrimaryY:  results[masterDisplayValueGroupBlue].y,
		whitePointX:   results[masterDisplayValueGroupWhitepoint].x,
		whitePointY:   results[masterDisplayValueGroupWhitepoint].y,
		maxLuminance:  results[masterDisplayValueGroupLuminance].x,
		minLuminance:  results[masterDisplayValueGroupLuminance].y,
	}, nil
}

func numbersInString(str string) (int64, error) {
	return strconv.ParseInt(nonNumericRegex.ReplaceAllString(str, ""), 10, 64)
}

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
	type pt struct{ x, y int64 }
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
		redPrimaryX:   m["R"].x,
		redPrimaryY:   m["R"].y,
		greenPrimaryX: m["G"].x,
		greenPrimaryY: m["G"].y,
		bluePrimaryX:  m["B"].x,
		bluePrimaryY:  m["B"].y,
		whitePointX:   m["WP"].x,
		whitePointY:   m["WP"].y,
		maxLuminance:  m["L"].x,
		minLuminance:  m["L"].y,
	}, nil
}
