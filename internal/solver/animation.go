package solver

import (
	"image"
	plt "image/color/palette"

	"golang.org/x/image/draw"
)

// countExplorablePixels scans the maze and counts the number of pixels that are not walls
func (s *Solver) countExplorablePixels() int {
	explorablePixels := 0
	for row := range s.maze.Bounds().Dy() {
		for col := range s.maze.Bounds().Dx() {
			if s.maze.RGBAAt(col, row) != s.palette.wall {
				explorablePixels++
			}
		}
	}
	return explorablePixels
}

// registerExploredPixels paints explored pixels, counts explored pixels so far and adds an animation frame accordingly
func (s *Solver) registerExploredPixels() {
	const totalExpectedFrames = 30

	explorablePixels := s.countExplorablePixels()
	pixelsExplored := 0

	for {
		select {
		case <-s.quit:
			return
		case pos := <-s.exploredPixels:
			s.maze.Set(pos.X, pos.Y, s.palette.explored)
			pixelsExplored++
			if pixelsExplored%(explorablePixels/totalExpectedFrames) == 0 {
				s.drawCurrentFrameToGIF()
			}
		}
	}
}

// drawCurrentFrameToGIF adds the current state of the maze as a frame of the animation.
func (s *Solver) drawCurrentFrameToGIF() {
	const (
		gifSize       = 500
		frameDuration = 20
	)

	// Create a paletted frame with the ratio as the inputImage
	frame := image.NewPaletted(image.Rect(0, 0, gifSize, gifSize*s.maze.Bounds().Dy()/s.maze.Bounds().Dx()), plt.Plan9)

	// Convert RGBA to paletted
	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, frameDuration)
}

// writeLastFrame write the final frame to the GIF with a highlighted solution and a longer duration
func (s *Solver) writeLastFrame() {
	stepsFromTreasure := s.solution

	// paint the path from entrance to treasure
	for stepsFromTreasure != nil {
		s.maze.Set(stepsFromTreasure.at.X, stepsFromTreasure.at.Y, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	const solutionFrameDuration = 300 // 3s
	s.drawCurrentFrameToGIF()
	s.animation.Delay[len(s.animation.Delay)-1] = solutionFrameDuration
}
