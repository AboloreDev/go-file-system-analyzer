package scanner

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/schollz/progressbar/v3"
)

// using concurrency to scan directories
// WalkDirectoryConcurrent walks directory using multiple workers

type FileInfo struct {
	Path string
	Size int64
	Hash string // MD5 hash for duplicate detection
}

func WalkDirectoryConcurrent(root string, workers int, calculateHash bool) ([]FileInfo, error) {
	 // Count files first
    var totalFiles int
    filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err == nil && !d.IsDir() {
            totalFiles++
        }
        return nil
    })
    
    bar := progressbar.Default(int64(totalFiles))
	// Channel to send file paths to workers
	// sending channels
	pathsChannel := make(chan string, 100)

	// Channel to collect results
	// receiving channels
	resultsChannel := make(chan FileInfo, 100)

	// channel to signal completion
	doneChannel := make(chan struct{})

	// Waitgroup for workers
	var waitgroup sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		waitgroup.Add(1)
		go worker(pathsChannel, resultsChannel, &waitgroup, calculateHash)
	}

	// close the go toutines
	go func() {
		waitgroup.Wait()
		close(resultsChannel)
		close(doneChannel)
	}()

	// walk direcoty and send pagth to workers
	go func() {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("Error accessing %s: %v\n", path, err)
				return nil

			}

			if !d.IsDir() {
				pathsChannel <- path
			}
			return nil
		})
		close(pathsChannel) // Signal no more paths
	}()

	// Collect results
	var files []FileInfo
	for file := range resultsChannel {
		files = append(files, file)
		        bar.Add(1)

	}

	<-doneChannel // Wait for cleanup

	return files, nil

}
// worker processes file paths from channel
func worker(paths <-chan string, results chan <- FileInfo, waitGroup *sync.WaitGroup, calculateHash bool) {
	defer waitGroup.Done()
	
	for path := range paths {
		info, err := fs.Stat(fs.FS(os.DirFS("/")), path)
		if err != nil {
			fmt.Printf("Error getting info for %s: %v\n", path, err)
			continue
		}
	
		fileInfo := FileInfo{
			Path: path,
			Size: info.Size(),
		}

		// Calculate hash if requested
        if calculateHash {
            hash, err := hashFile(path)
            if err != nil {
                fmt.Printf("Error hashing %s: %v\n", path, err)
            } else {
                fileInfo.Hash = hash
            }
        }
        
        results <- fileInfo
    }
}

// hashFile calculates MD5 hash of file
func hashFile(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()
    
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    
    return fmt.Sprintf("%x", hash.Sum(nil)), nil
}