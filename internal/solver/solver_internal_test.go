package solver

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver_findEntrance_success(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		want      image.Point
	}{
		"middle": {
			inputPath: "testdata/maze10_10.png",
			want:      image.Point{X: 0, Y: 5},
		},
		"400 px": {
			inputPath: "testdata/maze400_400.png",
			want:      image.Point{X: 0, Y: 200},
		},
		"treasure near entrance": {
			inputPath: "testdata/maze10_exit.png",
			want:      image.Point{X: 0, Y: 5},
		},
		"entrance in a corner": {
			inputPath: "testdata/maze10_corner.png",
			want:      image.Point{X: 0, Y: 0},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s, err := New(tc.inputPath)
			require.NoError(t, err)

			got, err := s.findEntrance()
			require.NoError(t, err)

			assert.Equal(t, got, tc.want)
		})
	}
}

func TestSolver_findEntrance_error(t *testing.T) {
	tt := map[string]struct {
		inputPath string
	}{
		"no entrance": {
			inputPath: "testdata/maze100_no_entrance.png",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s, err := New(tc.inputPath)
			require.NoError(t, err)

			got, err := s.findEntrance()
			require.Error(t, err)

			assert.Equal(t, got, image.Point{})

		})
	}
}
