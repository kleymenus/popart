package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("source image could not be loaded: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("source image format could not be decoded: %w", err)
	}

	return img, nil
}

func saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(file)

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
