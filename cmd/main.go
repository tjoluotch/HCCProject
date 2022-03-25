package main

import "HCCProject/internal/service"

func main() {
	//fh := filehandler.NewFileHandler()
	//if err := fh.CalculateTrendDataPoints(); err != nil {
	//	log.Fatal(err)
	//}
	service.Start()
}
