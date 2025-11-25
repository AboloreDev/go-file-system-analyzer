package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AboloreDev/go-file-system-analyzer/analyzer"
	"github.com/AboloreDev/go-file-system-analyzer/scanner"
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
	

	files, err := scanner.WalkDirectoryConcurrent(*dirPath, *workers, *findDupes)
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d files in total\n", len(files))

	// Further analysis can be done here
	if *findDupes {
    duplicates := analyzer.FindDuplicates(files)
    analyzer.PrintDuplicates(duplicates)
}

}
