package process

import (
	"3sigmas-monitorVisualization/pkg/data"
	"3sigmas-monitorVisualization/pkg/reader"
	"log"
)

func FindParser(records [][]string, parsers []reader.Parser) (reader.Parser, []data.Measure) {
	for _, parser := range parsers {
		measures, err := parser.Parse(records)
		if err == nil {
			return parser, measures
		}
		log.Printf("Parser %s rejected due to error: %s\n", parser.Source(), err)
	}
	return nil, nil
}
