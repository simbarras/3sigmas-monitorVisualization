package reader

import (
	"encoding/csv"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
)

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}

func Read(filepath string) [][]string {

	// Read csv from filepath
	file, err := os.Open(filepath)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	defer closeFile(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	log.Printf("Read %d lines from %s\n", len(records), filepath)
	return records
}
