package reader

import (
	"errors"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
	"strconv"
	"strings"
	"time"
)

type SenseiveParser struct{}

func (s *SenseiveParser) Parse(records [][]string) ([]data.Measure, error) {
	measures := make([]data.Measure, 0)
	for _, record := range records {

		if len(record) != 6 {
			return nil, errors.New("invalid record")
		}

		d, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			return nil, err
		}

		v, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, err
		}

		t := 0.0
		if len(record) > 5 {
			t, err = strconv.ParseFloat(record[5], 64)
			if err != nil {
				return nil, err
			}
		}
		m := data.SenseiveMeasure{
			DateTime:    d,
			Value:       v,
			Temperature: t,
			Captor:      record[1],
			Sensor:      record[2],
		}
		measures = append(measures, m)
	}
	log.Printf("Parsed %d measures\n", len(measures))
	if len(measures) == 0 {
		return nil, errors.New("no measures")
	}
	return measures, nil
}

// Sample:  Geosud-Demo_rail_2023-09-06_14-05-53.csv
// project name: Geosud-Demo_rail
// split at _20
func (s *SenseiveParser) ExtractProject(filename string) string {
	return strings.Split(filename, "_20")[0]
}

func (s *SenseiveParser) Source() string {
	return "senseive"
}
