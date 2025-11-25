package analyzer

import (
	"fmt"

	"github.com/AboloreDev/go-file-system-analyzer/scanner"
)

func FindDuplicates(files []scanner.FileInfo) map[string][]string {
    hashMap := make(map[string][]string)
    
    for _, file := range files {
        if file.Hash != "" {
            hashMap[file.Hash] = append(hashMap[file.Hash], file.Path)
        }
    }
    
    // Keep only hashes with multiple files
    duplicates := make(map[string][]string)
    for hash, paths := range hashMap {
        if len(paths) > 1 {
            duplicates[hash] = paths
        }
    }
    
    return duplicates
}

func PrintDuplicates(duplicates map[string][]string) {
    if len(duplicates) == 0 {
        fmt.Println("\nNo duplicates found!")
        return
    }
    
    fmt.Printf("\n=== Duplicate Files ===\n")
    fmt.Printf("Found %d sets of duplicates:\n\n", len(duplicates))
    
    for hash, paths := range duplicates {
        fmt.Printf("Hash: %s (%d copies)\n", hash[:8], len(paths))
        for _, path := range paths {
            fmt.Printf("  - %s\n", path)
        }
        fmt.Println()
    }
}