package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

const (
	ErrOpeningFile       = Error("unable to open image file")
	ErrDecodingError     = Error("unable to load input image")
	ErrExpectedRGBAImage = Error("expected RGBA image")
)

// openMaze opens an RGBA PNG image from a path.
func openMaze(imagePath string) (*image.RGBA, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %w", ErrOpeningFile, imagePath, err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %w", ErrDecodingError, imagePath, err)
	}

	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("%w, got %T", ErrExpectedRGBAImage, img)
	}

	return rgbaImage, nil
}
