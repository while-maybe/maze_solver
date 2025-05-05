package solver

import "image"

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
	return nil
}
