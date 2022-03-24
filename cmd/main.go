package main

import (
	"HCCProject/internal/filehandler"
	"log"
)

//const (
//	FILE_NAME         = "BATHROOM_VISIT_DATA.csv"
//	bathroomValidator = "Bathroom"
//	dateSeparator     = ":"
//)

func main() {
	fh := filehandler.NewFileHandler()
	if err := fh.CalculateTrendDataPoints(); err != nil {
		log.Fatal(err)
	}
}
