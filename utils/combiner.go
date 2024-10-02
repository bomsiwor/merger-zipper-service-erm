package utils

import (
	"errors"
	"fmt"
	"mymodule/entity"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func CombinePdf(outputFilename *string, inputFiles entity.Documents, workdir, outDir string) (string, error) {
	// Generate output file name
	filePrefix := strconv.Itoa(int(time.Now().Unix()))
	fileFormat := ".pdf"
	defaultFileName := filePrefix

	// Use system genreated file name if caller not specify the output filename
	if *outputFilename == "" {
		outputFilename = &defaultFileName
	}

	// Give extension
	*outputFilename = fmt.Sprintf("%s-%s%s", filePrefix, *outputFilename, fileFormat)

	// Combine require minimum 1 file
	if len(inputFiles) < 1 {
		return "", errors.New("not enough file")
	}

	// Final input file variable
	var finalInputFile []string

	// Clean the original path to create proper file path
	for _, input := range inputFiles {
		cleanPath := filepath.Clean(fmt.Sprintf("%s/%s", workdir, input.Path))

		// Final path
		finalInputFile = append(finalInputFile, filepath.FromSlash(cleanPath))
	}

	// Check if all file is PDF, if isn't convert it to PDF
	// temp file will deleted merging proces already done
	finalInputFile, tempFiles, err := convertAllToPdf(finalInputFile, workdir)
	if err != nil {
		return "", err
	}

	// Construct path by home and output file name
	cleanPath := filepath.Clean(fmt.Sprintf("%s/%s/%s", workdir, outDir, *outputFilename))

	// Final path
	finalPath := filepath.FromSlash(cleanPath)

	// Model for aggregate PDF
	model := model.NewDefaultConfiguration()

	err = api.MergeCreateFile(finalInputFile, finalPath, false, model)
	if err != nil {
		return "", err
	}

	// Delete file after merging process completed
	DeleteFiles(tempFiles)

	return fmt.Sprintf("%s/%s", outDir, *outputFilename), nil
}

// This function check if all data if pdf.
// If there is an image, convert it to pdf and store it on temporary folder.
// Then push path to finalPath. Return converted path, so the function caller can delete the temporary file
func convertAllToPdf(src []string, superFolder string) ([]string, []string, error) {
	// Variable to store temporary file path
	tempFiles := []string{}

	// Variable to store final path
	finalPath := []string{}

	// Loop through the source paths, check for PDF
	// PDF file pushed to finalPath
	for _, source := range src {
		// Check pdf file
		if filepath.Ext(source) == ".pdf" {
			finalPath = append(finalPath, source)
			continue
		}

		// COnvert image to pdf
		// Generate temp path filename
		wd, _ := os.Getwd()
		tempFilePath := GetFileNameFromPath(source) + ".pdf"
		tempPath := filepath.Join(wd, "doc", superFolder, "temp", tempFilePath)

		// Start converting
		err := ConvertImageToPdf([]string{source}, tempPath)
		if err != nil {
			return finalPath, tempFiles, err
		}

		// Push converted path to tempFile and finalPath
		tempFiles = append(tempFiles, tempPath)
		finalPath = append(finalPath, tempPath)
	}

	return finalPath, tempFiles, nil
}
