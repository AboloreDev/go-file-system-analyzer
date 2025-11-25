package analyzer

// import (
// 	"fmt"
// 	"sort"

// 	"github.com/AboloreDev/go-file-system-analyzer/scanner"
// )

// type Statistics struct {
// 	TotalFiles int
// 	TotalSize int64
// 	LargestFiles []scanner.FileInfo
// }

// func AnalyzeFiles(files []scanner.FileInfo, topN int) Statistics {
// 	var totalSize int64

// 	// Calculate total size
// 	for _, file := range files {
// 		totalSize += file.Size
// 	}

// 	// Sort by size(descending)
// 	sorted := make ([]scanner.FileInfo, len(files))

// 	copy(sorted, files)

// 	sort.Slice(sorted, func(i, j int) bool {
// 		return sorted[i].Size > sorted[j].Size
// 	})

// 	// Get top N largest files
// 	largest := sorted
// 	if len(sorted) > topN {
// 		largest = sorted[:topN]

// 	}

// 	return  Statistics{
// 		TotalFiles: len(files),
// 		TotalSize: totalSize,
// 		LargestFiles: largest,
// 	}
// }

// func (s Statistics) Print() {
// 	fmt.Printf("\n=== Analysis Results ===\n")
//     fmt.Printf("Total Files: %d\n", s.TotalFiles)
//     fmt.Printf("Total Size: %s\n", formatBytes(s.TotalSize))
//     fmt.Printf("\nTop %d Largest Files:\n", len(s.LargestFiles))

//     for i, file := range s.LargestFiles {
//         fmt.Printf("%d. %s - %s\n", i+1, file.Path, formatBytes(file.Size))
//     }

// }

// func formatBytes(bytes int64) string {
//     const unit = 1024
//     if bytes < unit {
//         return fmt.Sprintf("%d B", bytes)
//     }

//     div, exp := int64(unit), 0
//     for n := bytes / unit; n >= unit; n /= unit {
//         div *= unit
//         exp++
//     }

//     return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
// }