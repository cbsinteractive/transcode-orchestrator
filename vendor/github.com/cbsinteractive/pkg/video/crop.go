package video

// Crop holds offsets for top, bottom, left and right cropping, values are in pixels
type Crop struct {
	Top    int `json:"top,omitempty"`
	Bottom int `json:"bottom,omitempty"`
	Left   int `json:"left,omitempty"`
	Right  int `json:"right,omitempty"`
}

func (c Crop) Empty() bool {
	return c != (Crop{})
}
