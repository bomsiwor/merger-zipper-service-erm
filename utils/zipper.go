package utils

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mymodule/entity"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func CreateZip(outputFilename *string, inputFiles entity.Documents, workdir, outDir string) (string, error) {
	// Generate defalt file name
	filePrefix := strconv.Itoa(int(time.Now().Unix()))
	fileFormat := ".zip"
	defaultFileName := filePrefix

	if *outputFilename == "" {
		outputFilename = &defaultFileName
	}

	// Give extension
	*outputFilename = fmt.Sprintf("%s-%s%s", filePrefix, *outputFilename, fileFormat)

	// Create file
	// zipFile, err := os.Create(*outputFilename)
	// if err != nil {
	// 	panic(err)
	// }
	// defer zipFile.Close()

	// Use buffer
	zipFile := new(bytes.Buffer)

	// Create zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, input := range inputFiles {
		// Open file
		sourcePath := fmt.Sprintf("%s/%s", workdir, input.Path)

		file, err := os.Open(sourcePath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		// Get fileext
		ext := filepath.Ext(input.Path)

		f, err := zipWriter.Create(input.Name + "." + ext)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return "", err
		}
	}

	zipWriter.Close()

	// Construct path by home and output file name
	cleanPath := filepath.Clean(fmt.Sprintf("%s/%s/%s", workdir, outDir, *outputFilename))

	// Final path
	finalPath := filepath.FromSlash(cleanPath)
	// fmt.Println(finalPath)

	// Store buffer to file
	finalFile := os.WriteFile(finalPath, zipFile.Bytes(), 0644)
	if finalFile != nil {
		return "", errors.New("final file cannot be created")
	}

	return fmt.Sprintf("%s/%s", outDir, *outputFilename), nil
}
