package reader

import (
	"testing"
)

func TestSenseiveParser_Parse(t *testing.T) {

	p := SenseiveParser{}

	records := [][]string{
		{""},
	}

	_, err := p.Parse(records)

	if err == nil {
		t.Errorf("Empty records should return an error")
	}
}
