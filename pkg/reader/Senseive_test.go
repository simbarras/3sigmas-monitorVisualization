package reader

import (
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
	"testing"
	"time"
)

func parseErrorSenseive(records [][]string) bool {
	p := SenseiveParser{}
	_, err := p.Parse(records)
	if err == nil {
		return true
	}
	return false
}

func parseValidSenseive(records [][]string, expected []data.Measure) bool {
	p := SenseiveParser{}
	m, _ := p.Parse(records)
	if len(m) != len(expected) {
		return false
	}
	for i := range m {
		if m[i] != expected[i] {
			return false
		}
	}
	return true
}

func TestSenseiveParser_Parse_empty(t *testing.T) {
	records := [][]string{
		nil,
	}
	if parseErrorSenseive(records) {
		t.Errorf("Nil records should return an error")
	}

	records = [][]string{
		{},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Empty records should return an error")
	}

	records = [][]string{
		{""},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Record with empty string should return an error")
	}

	records = nil
	if parseErrorSenseive(records) {
		t.Errorf("Nil records should return an error")
	}
}

func TestSenseiveParser_Parse_invalid(t *testing.T) {
	records := [][]string{
		{"asdfas", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", "24.19"},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Invalid date should return an error")
	}

	records = [][]string{
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "asdfasdf", "24.19"},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Invalid value should return an error")
	}

	records = [][]string{
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "", "24.19"},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Empty value should return an error")
	}

	records = [][]string{
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", "asdfasdf"},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Invalid temperature should return an error")
	}

	records = [][]string{
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", ""},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Empty temperature should return an error")
	}
}

func TestSenseiveParser_Parse_valid(t *testing.T) {
	records := [][]string{
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", "24.19"},
	}
	d, _ := time.Parse("2006-01-02 15:04:05", "2023-09-06 13:32:00")
	expected := []data.Measure{
		data.SenseiveMeasure{
			DateTime:    d,
			Value:       4.203331,
			Temperature: 24.19,
			Captor:      "g_0_00-0_35",
			Sensor:      "Y Axis Beam Displacement",
		},
	}
	if !parseValidSenseive(records, expected) {
		t.Errorf("Valid record should not return an error")
	}
}

func TestSenseiveParser_Parse_trimble(t *testing.T) {
	records := [][]string{
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "520.0005", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorSenseive(records) {
		t.Errorf("Trimbe record should return an error")
	}
}

func TestSenseiveParser_Source(t *testing.T) {
	p := SenseiveParser{}
	if p.Source() != "senseive" {
		t.Errorf("Source should be senseive")
	}
}

func TestSenseiveParser_ExtractProject(t *testing.T) {
	p := SenseiveParser{}
	filename := "Geosud-Demo_rail_2023-09-06_13-52-46.csv"
	if p.ExtractProject(filename) != "Geosud-Demo_rail" {
		t.Errorf("Project name should be Geosud-Demo_rail")
	}
}

func FuzzSenseiveParser_ExtractProject(f *testing.F) {
	f.Add("Geosud-Demo_rail_2023-09-06_13-52-46.csv")
	f.Add("Geosud-Demo_rail_2023-09-06_13-52-46")
	f.Add("asdfasd09234ld dsf")
	f.Fuzz(func(t *testing.T, filename string) {
		p := SenseiveParser{}
		res := p.ExtractProject(filename)
		log.Printf("Name %s for filename %s\n", res, filename)
	})
}
