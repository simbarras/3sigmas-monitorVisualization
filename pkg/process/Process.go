package process

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/listener"
	"3sigmas-monitorVisualization/pkg/reader"
	"3sigmas-monitorVisualization/pkg/storer"
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

func Process(ftpListener *listener.FtpListener, influxStorer *storer.InfluxStorer, parser reader.Parser) {
	for {
		filepath := ftpListener.Listen()
		log.Printf("Process file %s\n", filepath)
		measures, err := parser.Parse(reader.ReadAndDelete(pkg.FtpLocalPath + "/" + filepath))
		if err != nil {
			log.Printf("File %s rejected due to error: %s\n", pkg.FtpLocalPath+"/"+filepath, err)
			ftpListener.RegisterBlacklist(filepath)
			continue
		}
		ftpListener.DeleteFile(filepath)
		go influxStorer.Store(parser.ExtractProject(filepath), parser.Source(), measures)
	}
}
