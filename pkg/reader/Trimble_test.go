package reader

import (
	"3sigmas-monitorVisualization/pkg/data"
	"log"
	"testing"
	"time"
)

func parseErrorTrimble(records [][]string) bool {
	p := TrimbleParser{}
	_, err := p.Parse(records)
	if err == nil {
		return true
	}
	return false
}

func parseValidTrimble(records [][]string, expected []data.Measure) (bool, []data.Measure) {
	p := TrimbleParser{}
	m, _ := p.Parse(records)
	if len(m) != len(expected) {
		return false, m
	}
	for i := range m {
		if m[i] != expected[i] {
			return false, m
		}
	}
	return true, m
}

func TestTrimbleParser_Parse_empty(t *testing.T) {
	records := [][]string{
		nil, nil,
	}
	if parseErrorTrimble(records) {
		t.Errorf("Nil records should return an error")
	}

	records = [][]string{
		{}, {},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Empty records should return an error")
	}

	records = [][]string{
		{""}, {""},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with empty string should return an error")
	}

	records = nil
	if parseErrorTrimble(records) {
		t.Errorf("Nil records should return an error")
	}
}

func TestTrimbleParser_Parse_invalid(t *testing.T) {
	records := [][]string{
		{},
		{"2gsdfg4", "st001", "1185070.0001", "2564880.0017", "520.0005", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid data should return an error")
	}

	records = [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "asldkfja", "2564880.0017", "520.0005", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid number at index 2 should return an error")
	}

	records = [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "asdfasfa", "520.0005", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid number at index 3 should return an error")
	}

	records = [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "asdfasdf", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid number at index 4 should return an error")
	}

	records = [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "520.0005", "asdfasdf", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid number at index 5 should return an error")
	}

	records = [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "520.0005", "1185070.0000", "asdfasdf", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid number at index 6 should return an error")
	}

	records = [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "520.0005", "1185070.0000", "2564880.0000", "asdfasdf", "0.0004566", "0.0002911", "0.0001465"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Record with invalid number at index 7 should return an error")
	}
}

func TestTrimbleParser_Parse_valid(t *testing.T) {
	records := [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "520.0005", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	d, _ := time.Parse("2006-01-02 15:04:05", "2023-09-13 08:44:12")
	expected := []data.Measure{
		data.TrimbleMeasure{
			DateTime:           d,
			Northing:           1185070.0001,
			Easting:            2564880.0017,
			Elevation:          520.0005,
			ReferenceNorthing:  1185070.0000,
			ReferenceEasting:   2564880.0000,
			ReferenceElevation: 520.0000,
			Captor:             "st001",
		},
	}
	res, resMeasure := parseValidTrimble(records, expected)
	if res {
		t.Errorf("Valid record should not return an error\n%v\n%v", res, resMeasure)
	}
}

func TestTrimbleParser_Parse_senseive(t *testing.T) {
	records := [][]string{
		{},
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", "24.19"},
	}
	if parseErrorTrimble(records) {
		t.Errorf("Senseive record should return an error")
	}
}

func TestTrimbleParser_Source(t *testing.T) {
	p := TrimbleParser{}
	if p.Source() != "trimble" {
		t.Errorf("Source should be trimble")
	}
}

func TestTrimbleParser_ExtractProject(t *testing.T) {
	p := TrimbleParser{}
	fileName := "Integrity Monitor [3s_230913]_20230913_090912_UTC.csv"
	if p.ExtractProject(fileName) != "3s_230913" {
		t.Errorf("Project name should be 3s_230913")
	}
}

func FuzzTrimbleParser_ExtractProject(f *testing.F) {
	f.Add("Integrity Monitor [3s_230913]_20230913_090912_UTC.csv")
	f.Add("Integrity Monitor [3s_230913]_20230913_090912_UTC")
	f.Add("asdlf;kjasldfjasdf")
	f.Fuzz(func(t *testing.T, filename string) {
		p := TrimbleParser{}
		res := p.ExtractProject(filename)
		log.Printf("Name %s for filename %s\n", res, filename)
	})
}
