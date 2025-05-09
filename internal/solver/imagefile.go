package solver

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log"
	"os"
	"strings"
)

const (
	ErrOpeningFile       = Error("unable to open image file")
	ErrDecodingError     = Error("unable to load input image")
	ErrExpectedRGBAImage = Error("expected RGBA image")
	ErrCreatingFile      = Error("unable to create output image file")
	ErrClosingFile       = Error("unable to close image file")
	ErrWritingFile       = Error("unable to write image file")
	ErrCreatingAnimGIF   = Error("unable to create output animated GIF")
	ErrEncodingGIF       = Error("unable to encode animated GIF")
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

// SaveSolution saves the image as a PNG file with the solution path highlighted.
func (s *Solver) SaveSolution(outputPath string) (err error) {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("%w at: %s", ErrCreatingFile, outputPath)
	}

	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.Join(err, ErrClosingFile, closeErr)
		}
	}()

	stepsFromTreasure := s.solution

	// paint the path from the last position back to the entrance
	for stepsFromTreasure != nil {
		s.maze.Set(stepsFromTreasure.at.X, stepsFromTreasure.at.Y, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	err = png.Encode(f, s.maze)
	if err != nil {
		return fmt.Errorf("%w at %s: %w", ErrWritingFile, outputPath, err)
	}

	gifPath := strings.Replace(outputPath, "png", "gif", -1)
	err = s.saveAnimation(gifPath)
	if err != nil {
		return ErrCreatingAnimGIF
	}

	return nil
}

// saveAnimation writes the animated GIF file
func (s *Solver) saveAnimation(gifPath string) error {
	outputImage, err := os.Create(gifPath)
	if err != nil {
		return fmt.Errorf("%w at: %s", ErrCreatingFile, gifPath)
	}

	defer func() {
		if closeErr := outputImage.Close(); closeErr != nil {
			err = errors.Join(err, ErrClosingFile, closeErr)
		}
	}()

	log.Printf("Animation contains %d frames\n", len(s.animation.Image))
	err = gif.EncodeAll(outputImage, s.animation)
	if err != nil {
		return ErrEncodingGIF
	}

	return nil
}
