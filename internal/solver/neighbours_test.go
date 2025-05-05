package solver

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_neighbours(t *testing.T) {
	tt := map[string]struct {
		p    image.Point
		want []image.Point
	}{
		"0, 0": {
			p: image.Point{0, 0},
			want: []image.Point{
				{0, 1}, {0, -1}, {1, 0}, {-1, 0},
			},
		},
		"1, 1": {
			p: image.Point{1, 1},
			want: []image.Point{
				{1, 2}, {1, 0}, {2, 1}, {0, 1},
			},
		},
		"8, -6": {
			p: image.Point{8, -6},
			want: []image.Point{
				{8, -5}, {8, -7}, {9, -6}, {7, -6},
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, tc.want, neighbours(tc.p))
		})
	}
}
