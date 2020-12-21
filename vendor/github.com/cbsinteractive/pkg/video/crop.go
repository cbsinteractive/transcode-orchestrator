package video

import "image"

// Crop holds offsets for top, bottom, left and right cropping, values are in pixels
type Crop struct {
	Left   int `json:"left,omitempty" redis-hash:"left,omitempty"`
	Top    int `json:"top,omitempty" redis-hash:"top,omitempty"`
	Right  int `json:"right,omitempty" redis-hash:"right,omitempty"`
	Bottom int `json:"bottom,omitempty" redis-hash:"bottom,omitempty"`
}

// Rect returns a Go rectangle, given the original source dimensions
func (c Crop) Rect(src image.Rectangle) image.Rectangle {
	return image.Rect(
		src.Min.X+c.Left,
		src.Min.Y+c.Top,
		src.Max.X-c.Right,
		src.Max.Y-c.Bottom,
	).Intersect(src).Canon()
}

// From sets c to the result of the cropping operation, it is the
// inverse of c.Rect
func (c *Crop) From(src, crop image.Rectangle) {
	src = src.Canon()
	crop = src.Intersect(crop).Canon()
	c.Left = crop.Min.X - src.Min.X
	c.Top = crop.Min.Y - src.Min.Y
	c.Right = src.Max.X - crop.Max.X
	c.Bottom = src.Max.Y - crop.Max.Y
}

func (c Crop) Empty() bool {
	return c == (Crop{})
}
