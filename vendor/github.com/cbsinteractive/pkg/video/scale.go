package video

import "image"

// Scale takes the src and crop rectangles
// and outputs a center-weighted crop rectangle
// with a bounds equal to src's aspect ratio.
//
// See https://github.com/as/croptest for a visual
// representation
func Scale(src, crop image.Rectangle) image.Rectangle {
	return scale(src, crop)
}

// Aspect returns r's aspect ratio as a vector
//
// For the rectangle (0,0)-(1920,1080), Aspect returns
// the point (16, 9).
func Aspect(r image.Rectangle) (ar image.Point) {
	return aspect(r)
}

func scale(s, c image.Rectangle) image.Rectangle {
	c = c.Intersect(s) // fix bounds on bad crops

	// obtain the center of the crop and save it
	// then translate the crop to (0,0)
	cc := center(c)
	c = c.Sub(c.Min)

	// compute the source's aspect ratio
	// in integral units and compute the
	// number of units cropped off, rounding
	// up to the nearest whole
	ar := Aspect(s)
	d := delta(ar, s, c)

	// pick the largest cut on either the x or y axis
	u := max(d.X, d.Y)

	// set the new crop rectangle to the number of pixels
	// those units represent in the source
	c.Max.X = s.Max.X - u*ar.X
	c.Max.Y = s.Max.Y - u*ar.Y

	// translate the new crop to the center of where the
	// old one used to be
	return c.Add(cc.Sub(center(c)))
}

func aspect(r image.Rectangle) (ar image.Point) {
	s := r.Size()
	g := gcd(s.X, s.Y)
	if g != 0 {
		return s.Div(g)
	}
	return s
}

func center(r image.Rectangle) image.Point {
	return r.Min.Add(r.Size().Div(2))
}

// delta returns the difference between src and r
// in units of the aspect ratio ar, rounded up to the
// nearest whole unit
func delta(ar image.Point, src, r image.Rectangle) (units image.Point) {
	if ar.X == 0 || ar.Y == 0{
		return image.ZP
	}
	return image.Pt(
		((src.Dx() - r.Dx() + ar.X - 1) / ar.X),
		((src.Dy() - r.Dy() + ar.Y - 1) / ar.Y),
	)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}