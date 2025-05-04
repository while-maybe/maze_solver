package main

import (
	"fmt"
	"log"
	"mazesolver/internal/solver"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)

	_, err := solver.New(inputFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
	}
}

// usage displays the usage syntax when calling this utility
func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}
