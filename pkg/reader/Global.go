package reader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Read(filepath string) [][]string {

	// Read csv from filepath
	fmt.Println("Reading file: ", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading csv: ", err)
		return nil
	}
	return records
}
