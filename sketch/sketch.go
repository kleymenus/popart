package sketch

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	_ "image/jpeg"
	"math/rand"
)

type Sketch struct {
	Input
	source            image.Image
	dc                *gg.Context
	sourceWidth       int
	sourceHeight      int
	strokeSize        float64
	currentStrokeSize float64
}

type Input struct {
	DestWidth                int
	DestHeight               int
	StrokeRatio              float64
	StrokeReduction          float64
	StrokeJitter             int
	StrokeInversionThreshold float64
	CurrentAlpha             float64
	AlphaIncrease            float64
	MinEdgeCount             int
	MaxEdgeCount             int
	Randomizer               *rand.Rand
}

func NewSketch(source image.Image, input Input) *Sketch {
	s := &Sketch{Input: input}

	bounds := source.Bounds()
	s.sourceWidth, s.sourceHeight = bounds.Max.X, bounds.Max.Y
	s.currentStrokeSize = s.StrokeRatio * float64(s.DestWidth)
	s.strokeSize = s.currentStrokeSize

	canvas := gg.NewContext(s.DestWidth, s.DestHeight)
	canvas.SetColor(color.Black)
	canvas.DrawRectangle(0, 0, float64(s.DestWidth), float64(s.DestHeight))
	canvas.FillPreserve()

	s.source = source
	s.dc = canvas

	return s
}

func (s *Sketch) Update() {
	rndX, rndY := s.obtainSource()
	destX, destY := s.determineDestination(rndX, rndY)
	r, g, b := rgb255(s.source.At(int(rndX), int(rndY)))
	edges := s.MinEdgeCount + s.Randomizer.Intn(s.MaxEdgeCount-s.MinEdgeCount+1)

	s.dc.SetRGBA255(r, g, b, int(s.CurrentAlpha))
	s.dc.DrawRegularPolygon(edges, destX, destY, s.strokeSize, s.Randomizer.Float64())
	s.dc.FillPreserve()

	if s.strokeSize <= s.StrokeInversionThreshold*s.currentStrokeSize {
		if (r+g+b)/3 < 128 {
			s.dc.SetRGBA255(255, 255, 255, int(s.CurrentAlpha*2))
		} else {
			s.dc.SetRGBA255(0, 0, 0, int(s.CurrentAlpha*2))
		}
	}
	s.dc.Stroke()

	s.strokeSize -= s.StrokeReduction * s.strokeSize
	s.CurrentAlpha += s.AlphaIncrease
}

func (s *Sketch) Output() image.Image {
	return s.dc.Image()
}

func (s *Sketch) obtainSource() (x, y float64) {
	x = s.Randomizer.Float64() * float64(s.sourceWidth)
	y = s.Randomizer.Float64() * float64(s.sourceHeight)
	return
}

func (s *Sketch) determineDestination(x, y float64) (dx, dy float64) {
	dx = x * float64(s.DestWidth) / float64(s.sourceWidth)
	dx += float64(s.randRange())
	dy = y * float64(s.DestHeight) / float64(s.sourceHeight)
	dy += float64(s.randRange())
	return
}

func (s *Sketch) randRange() int {
	return -s.StrokeJitter + s.Randomizer.Intn(2*s.StrokeJitter)
}
