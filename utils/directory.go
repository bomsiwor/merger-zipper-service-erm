package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// CHeck if the directory is exist or not
// Return true if directory already exists
func DirectoryChecker(path string) error {
	// Get directory from file path
	dirName := filepath.Dir(path)

	// Check using os.stat
	_, err := os.Stat(dirName)

	// If the directory does not exists, create new
	if os.IsNotExist(err) {
		err = CreateDirectory(dirName)

		if err != nil {
			return err
		}
	}

	return nil
}

func CreateDirectory(dirName string) error {
	// Create directory with 0755 permission
	// Using mkdirAll to create directory within its parent
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return err
	}

	return nil
}

func GetFileNameFromPath(filePath string) string {
	// Get file path base
	fileNameWithExt := filepath.Base(filePath)

	fileName := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))

	return fileName
}

func DeleteFiles(paths []string) {
	// Iterate over paths to delte each file
	for _, path := range paths {
		// Use os remove to delete file
		// We ignore the error so other program piece not break
		os.Remove(path)
	}
}
