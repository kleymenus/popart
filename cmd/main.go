package main

import (
	"github.com/kleymenus/popart/sketch"
	"log"
	"math/rand"
	"os"
	"time"
)

const TotalCyclesCount = 5000

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Missing parameters, please provide a path for the source and resulting images.")
	}

	img, err := loadImage(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	destWidth := 2000
	s := sketch.NewSketch(img, sketch.Input{
		DestWidth:                destWidth,
		DestHeight:               2000,
		StrokeRatio:              0.75,
		StrokeReduction:          0.002,
		StrokeInversionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		CurrentAlpha:             0.1,
		AlphaIncrease:            0.06,
		MinEdgeCount:             3,
		MaxEdgeCount:             4,
		Randomizer:               rand.New(rand.NewSource(time.Now().Unix())),
	})

	for i := 0; i < TotalCyclesCount; i++ {
		s.Update()
	}

	err = saveImage(s.Output(), os.Args[2])
	if err != nil {
		log.Panicln(err)
	}
}
