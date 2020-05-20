package mediaconvert

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	masterDisplayValueGroupRed        = "red"
	masterDisplayValueGroupBlue       = "blue"
	masterDisplayValueGroupGreen      = "green"
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

func parseMasterDisplay(encoded string) (masterDisplay, error) {
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
