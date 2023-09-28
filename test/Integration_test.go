package test

import (
	"3sigmas-monitorVisualization/pkg/process"
	"3sigmas-monitorVisualization/pkg/reader"
	"io"
	"log"
	"os"
	"testing"
)

func listFiles() ([]os.DirEntry, error) {
	files, err := os.ReadDir("../resources")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Found %d file(s)\n", len(files))
	log.Printf("Files: %v\n", files)
	return files, nil
}

func copyFile(filename string) error {
	if _, err := os.Stat("../backup"); os.IsNotExist(err) {
		err := os.Mkdir("../backup", 0755)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	sourceFile, err := os.Open("../resources/" + filename)
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

	destFile, err := os.Create("../backup/" + filename)
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

func TestIntegration(t *testing.T) {
	parsers := make([]reader.Parser, 0)
	parsers = append(parsers, &reader.SenseiveParser{})
	parsers = append(parsers, &reader.TrimbleParser{})

	files, err := listFiles()
	if err != nil {
		t.Error(err)
		return
	}

	for _, file := range files {
		err = copyFile(file.Name())
		if err != nil {
			t.Error(err)
			return
		}
		parser, measures := process.FindParser(reader.ReadAndDelete("../backup/"+file.Name()), parsers)
		if parser == nil {
			if file.Name() != "Geosud-Demo_rail_2023-09-07_10-06-25-copy.csv" {
				t.Errorf("No parser found for file %s\n", file.Name())
			}
			continue
		}
		if len(measures) == 0 {
			t.Errorf("No measures found for file %s\n", file.Name())
			continue
		}
	}
}
