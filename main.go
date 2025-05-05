package main

import (
	"fmt"
	"log"
	"mazesolver/internal/solver"
	"os"
)

const (
	ErrBadUsage = Error("usage: mazesolver input.png output.png")
)

func main() {
	if len(os.Args) != 3 {
		exit(ErrBadUsage)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)

	s, err := solver.New(inputFile)
	if err != nil {
		exit(err)
	}

	err = s.Solve()
	if err != nil {
		exit(err)
	}

	err = s.SaveSolution(outputFile)
	if err != nil {
		exit(err)
	}
}

// exit prints the error and exits to OS
func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}
