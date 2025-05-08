package solver

import (
	"image"
	"log"
	"sync"
)

// explore one path and publish it to the s.pathsToExplore channel any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		// this is a safety net. It should be used, but when it's needed, at least it's there.
		return
	}

	pos := pathToBranch.at

	for {
		s.mutex.Lock()
		s.maze.SetRGBA(pos.X, pos.Y, s.palette.explored)
		s.mutex.Unlock()

		// is it time to quit? Did another goroutine found the treasure?
		select {
		case <-s.quit:
			return
		case s.exploredPixels <- pos:

			// continue exploring
		}

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
					// s.quit <- struct{}{}
					close(s.quit)
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
			select {
			// remember s.quit only returns a zero value after channel is closed!
			case <-s.quit:
				log.Printf("I'm an unlucky branch, someone else found the treasure. I gave up at %v", pos)
				return
			case s.pathsToExplore <- branch:
				//continue execution
			}
		}

		// continue exploring
		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}

}

// listenToBranches creates a new goroutine for each branch published in s.pathsToExplore
func (s *Solver) listenToBranches() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for {
		select {
		case <-s.quit:
			log.Println("Treasure found, stopping worker")
			return
		case p := <-s.pathsToExplore:
			wg.Add(1)

			go func(path *path) {
				defer wg.Done()
				s.explore(p)
			}(p)
		}
	}
}

// isPreviousStep returns true if the given point is the previous position of the path.
func (p path) isPreviousStep(n image.Point) bool {
	return p.previousStep != nil && p.previousStep.at == n
}
