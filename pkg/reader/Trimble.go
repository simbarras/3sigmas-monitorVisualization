package reader

import (
	"3sigmas-monitorVisualization/pkg/data"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type TrimbleParser struct{}

func (t *TrimbleParser) Parse(records [][]string) ([]data.Measure, error) {
	measures := make([]data.Measure, 0)
	head := true
	for _, record := range records {
		if head {
			head = false
			continue
		}
		if len(record) != 11 {
			return nil, errors.New("invalid record length")
		}
		d, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			return nil, err
		}
		n, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		ea, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, err
		}
		el, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, err
		}
		rn, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, err
		}
		rea, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			return nil, err
		}
		rel, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			return nil, err
		}
		m := data.TrimbleMeasure{
			DateTime:           d,
			Northing:           n,
			Easting:            ea,
			Elevation:          el,
			ReferenceNorthing:  rn,
			ReferenceEasting:   rea,
			ReferenceElevation: rel,
			Captor:             record[1],
		}
		measures = append(measures, m)
	}
	log.Printf("Parsed %d measures\n", len(measures))
	if len(measures) == 0 {
		return nil, errors.New("no measures")
	}
	return measures, nil
}

// Sample: Integrity Monitor [3s_230913]_20230913_113414_UTC
// project name: 3s_230913
// Extract between [ and ]
func (t *TrimbleParser) ExtractProject(filename string) string {
	start := strings.Index(filename, "[")
	if start == -1 {
		return ""
	}
	return filename[start+1 : strings.Index(filename, "]")]
}

func (t *TrimbleParser) Source() string {
	return "trimble"
}
