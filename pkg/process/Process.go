package process

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/data"
	"3sigmas-monitorVisualization/pkg/reader"
	"github.com/getsentry/sentry-go"
	"log"
)

func SetSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://ff671ea4c1f2ace1811282b96669e095@o4505048574001152.ingest.sentry.io/4505841145675776",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Release:          pkg.Version,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

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
