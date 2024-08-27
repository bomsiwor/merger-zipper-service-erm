package utils

import (
	"errors"
	"fmt"
	"mymodule/entity"
	"path/filepath"
	"strconv"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func CombinePdf(outputFilename *string, inputFiles entity.Documents, workdir, outDir string) (string, error) {
	filePrefix := strconv.Itoa(int(time.Now().Unix()))
	fileFormat := ".pdf"
	defaultFileName := filePrefix

	if *outputFilename == "" {
		outputFilename = &defaultFileName
	}

	// Give extension
	*outputFilename = fmt.Sprintf("%s-%s%s", filePrefix, *outputFilename, fileFormat)

	if len(inputFiles) < 2 {
		return "", errors.New("not enough file")
	}

	var finalInputFile []string

	for _, input := range inputFiles {
		cleanPath := filepath.Clean(fmt.Sprintf("%s/%s", workdir, input.Path))

		// Final path
		finalInputFile = append(finalInputFile, filepath.FromSlash(cleanPath))
	}

	// Construct path by home and output file name
	cleanPath := filepath.Clean(fmt.Sprintf("%s/%s/%s", workdir, outDir, *outputFilename))

	// Final path
	finalPath := filepath.FromSlash(cleanPath)

	// Model for aggregate PDF
	model := model.NewDefaultConfiguration()

	err := api.MergeCreateFile(finalInputFile, finalPath, false, model)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", outDir, *outputFilename), nil
}
