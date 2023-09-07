package reader

import (
	"3sigmas-monitorVisualization/pkg/reader/data"
	"fmt"
	"strconv"
	"time"
)

func SenseiveParse(records [][]string) []data.Measure {
	measures := make([]data.Measure, 0)
	for _, record := range records {

		d, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			panic(err)
		}

		v, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			panic(err)
		}

		t := 0.0
		if len(record) > 5 {
			t, err = strconv.ParseFloat(record[5], 64)
			if err != nil {
				panic(err)
			}
		}

		m := data.Measure{
			Date:        d,
			Value:       v,
			Temperature: t,
			Captor:      record[1],
			Sensor:      record[2],
		}
		fmt.Println("Created measure: ", m)
		measures = append(measures, m)
	}
	return measures
}
