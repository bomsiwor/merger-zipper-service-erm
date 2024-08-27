package utils

import (
	"errors"
	"os"
)

func WorkdirRouting(apiKey string) (string, error) {
	switch apiKey {
	case os.Getenv("PRODUCTION_API_KEY"):
		return "doc/production", nil

	case os.Getenv("STAGING_API_KEY"):
		return "doc/staging", nil

	default:
		return "", errors.New("source path not found")
	}
}