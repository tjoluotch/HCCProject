package main

import (
	"HCCProject/internal/filehandler"
	"log"
)

func main() {
	fh := filehandler.NewFileHandler()
	if err := fh.CalculateTrendDataPoints(); err != nil {
		log.Fatal(err)
	}
}
