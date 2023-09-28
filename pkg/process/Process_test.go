package process

import (
	"3sigmas-monitorVisualization/pkg/reader"
	"testing"
)

func TestSetSentry(t *testing.T) {
	SetSentry()
}

func TestFindParser(t *testing.T) {
	parsers := make([]reader.Parser, 0)
	parsers = append(parsers, &reader.SenseiveParser{})
	parsers = append(parsers, &reader.TrimbleParser{})

	records1 := [][]string{
		{"2023-09-06 13:32:00", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", "24.19"},
	}
	parser, _ := FindParser(records1, parsers)
	if parser != parsers[0] {
		t.Errorf("Expected %s, got %s", parsers[0].Source(), parser.Source())
	}

	records2 := [][]string{
		{},
		{"2023-09-13 08:44:12", "st001", "1185070.0001", "2564880.0017", "520.0005", "1185070.0000", "2564880.0000", "520.0000", "0.0004566", "0.0002911", "0.0001465"},
	}
	parser, _ = FindParser(records2, parsers)
	if parser != parsers[1] {
		t.Errorf("Expected %s, got %s", parsers[1].Source(), parser.Source())
	}

	records3 := [][]string{
		{"asdfasdf", "g_0_00-0_35", "Y Axis Beam Displacement", "mm", "4.203331", "asdfasdf"},
	}
	parser, _ = FindParser(records3, parsers)
	if parser != nil {
		t.Errorf("Expected nil, got %s", parser.Source())
	}
}
