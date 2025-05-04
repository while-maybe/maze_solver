package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenImage_errors(t *testing.T) {
	tt := map[string]struct {
		input string
		err   string
	}{
		"no such file": {
			input: "doesnt_exist.png",
			err:   "no such file or directory",
		},
		"not an rgba png": {
			input: "testdata/rgb.png",
			err:   "expected RGBA image, got *image.Paletted",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			img, err := openMaze(tc.input)

			assert.Nil(t, img)
			assert.Error(t, err)
			assert.ErrorContains(t, err, tc.err)
		})
	}
}
