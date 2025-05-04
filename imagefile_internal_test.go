package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenImage_errors(t *testing.T) {
	tt := map[string]struct {
		input string
		err   Error
	}{
		"no such file": {
			input: "doesnt_exist.png",
			err:   ErrOpeningFile,
		},
		"can't decode data": {
			input: "testdata/empty_file.png",
			err:   ErrDecodingError,
		},
		"not an rgba png": {
			input: "testdata/rgb.png",
			err:   ErrExpectedRGBAImage,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			img, err := openMaze(tc.input)

			assert.Nil(t, img)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.err)
			// assert.ErrorContains(t, err, tc.err)
		})
	}
}
