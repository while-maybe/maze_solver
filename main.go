package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving is as %q", inputFile, outputFile)

	_, err := openMaze(inputFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
	}
}

// usage displays the usage when calling this utility, returns an error
func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}
