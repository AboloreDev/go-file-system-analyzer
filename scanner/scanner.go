package scanner

// import (
// 	"fmt"
// 	"io/fs"
// 	"path/filepath"
// )

// // Fileinfo that holds the file maetadata
// type FileInfo struct {
// 	Path string
// 	Size int64
// }

// // Walkdirectory walks a directory and returns all files
// func WalkDirectoy(root string) ([]FileInfo, error) {
// 	var files []FileInfo

// 	err := filepath.WalkDir(root, func(path string, directory fs.DirEntry, err error) error {
// 		if err != nil {
// 			// log the error but continue walking
// 			fmt.Printf("Error accessing path %q: %v\n", path, err)
// 			return nil
// 		}

// 		// skip diretories/folder, only process files
// 		if directory.IsDir() {
// 			return nil
// 		}

// 		info, err := directory.Info()

// 		if err != nil {
// 			fmt.Printf("Error getting info for file %q: %v\n", path, err)
// 			return nil
// 		}

// 		files = append(files, FileInfo{
// 			Path: path,
// 			Size: info.Size(),
// 		})

// 		return nil
// 	})

// 	return files, err
// }
