package data

import (
	"testing"
	"time"
)

func TestSenseiveMeasure(t *testing.T) {
	d, _ := time.Parse("2006-01-02T15:04:05Z", "2019-01-01T00:00:00Z")
	m := SenseiveMeasure{
		DateTime:    d,
		Value:       1.0,
		Temperature: 2.0,
		Captor:      "captor",
		Sensor:      "sensor",
	}

	if m.String() != "2019-01-01T00:00:00Z captor sensor 1.000000 2.000000" {
		t.Errorf("String() failed (%s)", m.String())
	}
	if m.Tags()["type"] != "sensor" {
		t.Error("Tags() failed")
	}
	if m.Fields()["value"] != 1.0 {
		t.Error("Fields() failed")
	}
	if m.Measurement() != "captor" {
		t.Error("Measurement() failed")
	}
	if m.Date().Format("2006-01-02T15:04:05Z") != "2019-01-01T00:00:00Z" {
		t.Error("Date() failed")
	}
}

func TestTrimbleMeasure(t *testing.T) {
	d, _ := time.Parse("2006-01-02T15:04:05Z", "2019-01-01T00:00:00Z")
	m := TrimbleMeasure{
		DateTime:           d,
		Captor:             "captor",
		Northing:           1.0,
		Easting:            2.0,
		Elevation:          3.0,
		ReferenceNorthing:  4.0,
		ReferenceEasting:   5.0,
		ReferenceElevation: 6.0,
	}

	if m.String() != "2019-01-01T00:00:00Z 1.000000 2.000000 3.000000" {
		t.Errorf("String() failed (%s)", m.String())
	}
	if len(m.Tags()) != 0 {
		t.Error("Tags() failed")
	}
	if m.Fields()["northing"] != 1.0 {
		t.Error("Fields() failed")
	}
	if m.Fields()["easting"] != 2.0 {
		t.Error("Fields() failed")
	}
	if m.Fields()["elevation"] != 3.0 {
		t.Error("Fields() failed")
	}
	if m.Fields()["referenceNorthing"] != 4.0 {
		t.Error("Fields() failed")
	}
	if m.Fields()["referenceEasting"] != 5.0 {
		t.Error("Fields() failed")
	}
	if m.Fields()["referenceElevation"] != 6.0 {
		t.Error("Fields() failed")
	}
	if m.Measurement() != "captor" {
		t.Error("Measurement() failed")
	}
	if m.Date().Format("2006-01-02T15:04:05Z") != "2019-01-01T00:00:00Z" {
		t.Error("Date() failed")
	}
}
