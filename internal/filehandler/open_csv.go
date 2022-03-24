package filehandler

import (
	"errors"
	"os"
	"strings"
)

// OpenFile opens the CSV file and returns the *File with any error reporting
func OpenFile(file string) (*os.File, error) {
	// if file is incorrect format - not .csv return error
	if !strings.HasSuffix(file, fileType) {
		return nil, errors.New(incorrectFmtErr)
	}

	// open csv file
	visitsDataCSV, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return visitsDataCSV, nil
}
