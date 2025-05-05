package solver

import (
	"fmt"
	"image"
	"log"
)

const (
	ErrNoEntrance = Error("entrance position not found")
)

// Solver is capable of finding the path from the entrance to the treasure.
// The maze has to be an RGB image
type Solver struct {
	maze    *image.RGBA
	palette palette
}

// New builds a Solver by taking the path to the PNG maze, encoded in RGBA.
func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
		return nil, ErrOpeningFile
	}

	return &Solver{
		maze:    img,
		palette: defaultPalette(),
	}, nil
}

// Solve finds the path from the entrance to the treasure.
func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("solver didn't find an entrance: %w", err)
	}

	log.Printf("starting at pos: %v", entrance)
	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (image.Point, error) {
	for row := s.maze.Bounds().Min.Y; row < s.maze.Bounds().Max.Y; row++ {
		for col := s.maze.Bounds().Min.X; col < s.maze.Bounds().Max.X; col++ {
			if s.maze.RGBAAt(col, row) == s.palette.entrance {
				return image.Point{col, row}, nil
			}
		}
	}
	return image.Point{}, ErrNoEntrance
}
