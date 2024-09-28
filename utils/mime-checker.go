package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// getFileMimeType checks the MIME type of the file
func GetFileMimeType(file *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)

	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	file.Seek(0, io.SeekStart) // Reset file read pointer

	mimeType := http.DetectContentType(buffer)

	return mimeType, nil
}

// isValidMimeType checks if the MIME type is png, jpg or jpeg
func IsValidImageMime(mimeType string) bool {
	return mimeType == "image/png" || mimeType == "image/jpeg"
}

func IsValidPdfMime(mimeType string) bool {
	return mimeType == "application/pdf"
}
