package solver

import "image/color"

// palette contains the colours of the different types of pixels in our maze.
type palette struct {
	wall     color.RGBA
	path     color.RGBA
	entrance color.RGBA
	treasure color.RGBA
	solution color.RGBA
	explored color.RGBA
}

// defaultPalette returns the colour palette of our maze.
func defaultPalette() palette {
	return palette{
		wall:     color.RGBA{R: 0, G: 0, B: 0, A: 255},       // black
		path:     color.RGBA{R: 255, G: 255, B: 255, A: 255}, // white
		entrance: color.RGBA{R: 0, G: 191, B: 255, A: 255},   // deep sky blue
		treasure: color.RGBA{R: 255, G: 0, B: 128, A: 255},   // pink
		solution: color.RGBA{R: 255, G: 140, B: 0, A: 255},   // orange
		explored: color.RGBA{R: 0, G: 254, B: 254, A: 255},   // cyan
	}
}
