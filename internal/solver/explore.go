package solver

import (
	"image"
	"log"
)

// explore one path and publish it to the s.pathsToExplore channel any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		// this is a safety net. It should be used, but when it's needed, at least it's there.
		return
	}

	pos := pathToBranch.at

	for !s.solutionFound() {
		// we'll have up to 3 neighbours to explore.
		candidates := make([]image.Point, 0, 3)

		// look at all the neighbours for destination candidates
		for _, n := range neighbours(pos) {
			if pathToBranch.isPreviousStep(n) {
				// we don't want to go back
				continue
			}

			// we look at the color
			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.palette.treasure:
				s.mutex.Lock()
				defer s.mutex.Unlock()

				if s.solution == nil {

					s.solution = &path{previousStep: pathToBranch, at: n}
					log.Printf("Treasure found at %v", n)
				}

				return
			case s.palette.path:
				candidates = append(candidates, n)
			}
		}

		if len(candidates) == 0 {
			log.Printf("This is not the way: %v", pos)
			return
		}

		// notify the channel when there's more than one possible paths to explore
		for _, candidate := range candidates[1:] {
			branch := &path{previousStep: pathToBranch, at: candidate}
			s.pathsToExplore <- branch
		}

		// continue exploring
		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}

}

// listenToBranches creates a new goroutine for each branch published in s.pathsToExplore
func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		go s.explore(p)

		if s.solutionFound() {
			return
		}
	}
}

// isPreviousStep returns true if the given point is the previous position of the path.
func (p path) isPreviousStep(n image.Point) bool {
	return p.previousStep != nil && p.previousStep.at == n
}

// solutionFound returns true if the solution has been found
func (s *Solver) solutionFound() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.solution != nil
}
