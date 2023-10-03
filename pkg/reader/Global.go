package reader

import (
	"encoding/csv"
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
	"os"
)

type Parser interface {
	Parse(records [][]string) ([]data.Measure, error)
	ExtractProject(filename string) string
	Source() string
}

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
		return nil
	}
	defer closeAndDelete(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	log.Printf("Read %d lines from %s\n", len(records), filepath)
	return records
}
