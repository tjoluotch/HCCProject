package filehandler

import "os"

// openFile opens the CSV file and returns the *File with any error reporting
func openFile() (*os.File, error) {
	// open csv file
	visitsDataCSV, err := os.Open(FILE_NAME)
	if err != nil {
		return nil, err
	}
	return visitsDataCSV, nil
}
