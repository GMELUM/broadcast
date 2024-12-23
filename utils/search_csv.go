package utils

import (
	"errors"
	"path/filepath"
)

func SearchCSV() (string, error) {

	matches, err := filepath.Glob(filepath.Join(".", "*.csv"))
	if err != nil {
		return "", errors.New("file search error")
	}
	if len(matches) == 0 {
		return "", errors.New("csv file not found")
	}
	if len(matches) > 1 {
		return "", errors.New("more than one csv file found")
	}

	return matches[0], nil

}
