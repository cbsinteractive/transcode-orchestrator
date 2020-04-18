package timecode

import (
	"fmt"
	"time"
)

// Range is a pair of decimal seconds defining a time interval
// starting at Range[0] and ending at Range[1]
type Range [2]float64

// Canon returns the range in proper order, where r[0] <= r[1]
func (r Range) Canon() Range {
	if r[0] > r[1] {
		r[0], r[1] = r[1], r[0]
	}
	return r
}

// Size returns the duration of the Range
func (r Range) Size() time.Duration {
	dx := r[1] - r[0]
	if dx < 0 {
		dx = -dx
	}
	return time.Duration(dx * float64(time.Second))
}

func (r Range) String() string {
	const s = float64(time.Second)
	return fmt.Sprintf("(%s-%s)", time.Duration(r[0]*s), time.Duration(r[1]*s))
}

// Timecode outputs the timecode in HH:MM:SS:FF format
func (r Range) Timecode(fps float64) string {
	return toString(r[1], fps)
}

// Timecodes outputs the start and end timecodes in HH:MM:SS:FF format
// TODO(as): should replace Timecode with this. The Range float64 types
// might be better off as custom duration types that use integral units, rather
// than float64, we could add methods on these directly and then remove
// Timecode
func (r Range) Timecodes(fps float64) (string, string) {
	return toString(r[0], fps), toString(r[1], fps)
}

func (r Range) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%f,%f]", r[0], r[1])), nil
}

// Parse parses an input string in HH:MM:SS:FF, HH:MM:SS;FF, or
// HH:MM:SS format, defined by the following convention
// HH = hour, MM = minute, SS = second, and FF is the frame number, the
// frameRate argument is either 0, or a fractional frame rate upon which
// to calculate the precise Range value based on the FF argument, if present.
func Parse(timecode string, fps float64) (Range, error) {
	if fps == 0 {
		fps = defaultFps
	}
	var (
		h, m, s float64
		f       uint64
	)
	n, err := fmt.Sscanf(timecode, "%f:%f:%f:%d", &h, &m, &s, &f)
	if n < 3 {
		n, _ = fmt.Sscanf(timecode, "%f:%f:%f;%d", &h, &m, &s, &f)
	}
	if n < 3 {
		return Range{}, err
	}

	// To avoid floating point issues, we convert frames per second
	// into nanosecond per frame. If we have 1 fps, then frame exposure
	// time takes 1e9 nanoseconds, for 2 fps, 1e9/2 ns, and so forth. We
	// can then multiply the frame number by the number of nanoseconds
	// per frame and convert that into a floating point representation
	// as the final step
	nspf := uint64(1e9 / fps)                 // ns/frame
	fdur := time.Duration(nspf * f).Seconds() // n.o. seconds these frames take up
	return Range{0, h*3600 + m*60 + s + fdur}, err
}

// toString converts the float64 number of seconds s into a string timecode
func toString(s float64, fps float64) string {
	if fps == 0 {
		fps = defaultFps
	}
	{
		d := int64(s)
		h := d / 3600
		d %= 3600
		m := d / 60
		m %= 60
		s := d % 60
		f := 0 // TODO(as): frame number
		return fmt.Sprintf("%02d:%02d:%02d:%02d", h, m, s, f)
	}
}
