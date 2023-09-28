package reader

import (
	"io"
	"log"
	"os"
	"testing"
)

func copyFile(filename string) error {
	// Create backup folder if it does not exist
	if _, err := os.Stat("../../backup"); os.IsNotExist(err) {
		err := os.Mkdir("../../backup", 0755)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	// List files in ./resources
	files, err := os.ReadDir("../../resources")
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Found %d file(s)\n", len(files))
	log.Printf("Files: %v\n", files)

	sourceFile, err := os.Open("../../resources/" + filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(sourceFile)
	destFile, err := os.Create("../../backup/" + filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(destFile)
	nbBytes, err := io.Copy(destFile, sourceFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Copied %d bytes from %s to %s\n", nbBytes, sourceFile.Name(), destFile.Name())
	return nil
}

func TestReadAndDelete(t *testing.T) {
	const filename = "Geosud-Demo_rail_2023-09-07_10-06-25.csv"
	// Copy files from ./resources to ./backup
	err := copyFile(filename)
	if err != nil {
		t.Errorf("copyFile() = %s; want nil", err)
		return
	}

	// ReadAndDelete from ./backup
	records := ReadAndDelete("../../backup/" + filename)
	if len(records) != 1221 {
		t.Errorf("ReadAndDelete() = %d; want 1221", len(records))
	}

	records = ReadAndDelete("../../backup/" + filename)
	if len(records) != 0 {
		t.Errorf("ReadAndDelete() = %d; want 0", len(records))
	}

	// Check if the files are deleted
	_, err = os.Stat("./backup/" + filename)
	if !os.IsNotExist(err) {
		t.Errorf("ReadAndDelete() = %s; want %s", err, "file does not exist")
	}
}

func TestReadAndDelete_empty(t *testing.T) {
	const filename = "Geosud-Demo_rail_2023-09-07_10-06-25-copy.csv"
	// Copy files from ./resources to ./backup
	err := copyFile(filename)
	if err != nil {
		t.Errorf("copyFile() = %s; want nil", err)
		return
	}

	// ReadAndDelete from ./backup
	records := ReadAndDelete("../../backup/" + filename)
	if len(records) != 0 {
		t.Errorf("ReadAndDelete() = %d; want 1221", len(records))
	}

	records = ReadAndDelete("../../backup/" + filename)
	if len(records) != 0 {
		t.Errorf("ReadAndDelete() = %d; want 0", len(records))
	}

	// Check if the files are deleted
	_, err = os.Stat("./backup/" + filename)
	if !os.IsNotExist(err) {
		t.Errorf("ReadAndDelete() = %s; want %s", err, "file does not exist")
	}
}
