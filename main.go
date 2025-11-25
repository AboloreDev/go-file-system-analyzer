package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"os"
)

func main() {
	// flags definition
	dirPath := flag.String("path", ".", "Path to the directory to analyze")
	workers := flag.Int("workers", 4, "Number of concurrent workers")
	findDupes := flag.Bool("duplicates", false, "find duplicate files")

	flag.Parse()

	// File path validation
	if _, err := os.Stat(*dirPath); os.IsNotExist(err) {
		fmt.Printf("Error: The specified path does not exist: %s\n", *dirPath)
		os.Exit(1)
	}

	fmt.Printf("Analyzing directory: %s\n", *dirPath)
	fmt.Printf("Using %d workers\n", *workers)

	// Using the duplicates flag
	if *findDupes {
		fmt.Println("Duplicate file detection is enabled")
	} else {
		fmt.Println("Duplicate file detection is disabled")
	}

	files, err := scanner.Walkdirectory(*dirPath)
}
