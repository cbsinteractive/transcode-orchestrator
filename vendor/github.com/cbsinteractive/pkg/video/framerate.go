package video

// Framerate contains individual integer elements of a fractional framerate
// including both the numerator and its divisor
type Framerate struct {
	Numerator   int `json:"numerator,omitempty" redis-hash:"numerator,omitempty"`
	Denominator int `json:"denominator,omitempty" redis-hash:"denominator,omitempty"`
}

func (f Framerate) Empty() bool {
	return f.Denominator == 0 || f.Numerator == 0
}
