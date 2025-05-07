package solver

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver_explorer(t *testing.T) {
	tt := map[string]struct {
		inputImage string
		wantSize   int
	}{
		"cross": {
			inputImage: "testdata/explore_cross.png",
			wantSize:   2,
		},
		"dead end": {
			inputImage: "testdata/explore_deadend.png",
			wantSize:   0,
		},
		"double": {
			inputImage: "testdata/explore_double.png",
			wantSize:   1,
		},
		"treasure": {
			inputImage: "testdata/explore_treasure.png",
			wantSize:   0,
		},
		"treasure only": {
			inputImage: "testdata/explore_treasureonly.png",
			wantSize:   0,
		},
	}

	for name, tc := range tt {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			maze, err := openMaze(tc.inputImage)
			require.NoError(t, err)

			s := &Solver{
				maze:           maze,
				palette:        defaultPalette(),
				pathsToExplore: make(chan *path, 3),
			}

			s.explore(&path{at: image.Point{X: 0, Y: 2}})

			assert.Equal(t, tc.wantSize, len(s.pathsToExplore))
		})
	}
}
