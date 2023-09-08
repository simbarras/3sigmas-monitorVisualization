package reader

import (
	"encoding/csv"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
)

func closeAndDelete(file *os.File) {
	err := file.Close()
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	err = os.Remove(file.Name())
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}

func ReadAndDelete(filepath string) [][]string {

	// ReadAndDelete csv from filepath
	file, err := os.Open(filepath)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	defer closeAndDelete(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	log.Printf("Read %d lines from %s\n", len(records), filepath)
	return records
}
