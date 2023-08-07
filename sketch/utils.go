package sketch

import "image/color"

func rgb255(c color.Color) (r, g, b int) {
	r0, g0, b0, _ := c.RGBA()
	return int(r0 / 257), int(g0 / 257), int(b0 / 257)
}

func (s *Sketch) randRange() int {
	return -s.StrokeJitter + s.Randomizer.Intn(2*s.StrokeJitter)
}
