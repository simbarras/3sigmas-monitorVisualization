package main

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/data"
	"3sigmas-monitorVisualization/pkg/listener"
	"3sigmas-monitorVisualization/pkg/process"
	"3sigmas-monitorVisualization/pkg/reader"
	"3sigmas-monitorVisualization/pkg/storer"
	"github.com/getsentry/sentry-go"
	"log"
)

func main() {
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

	log.Printf("App started in release %s\n", pkg.Version)

	env := data.ReadEnv()
	log.Printf("Launching senseive reader with env: \n%+v\n", env)

	influxStorer := storer.NewInfluxStorer(env)
	ftpListener := listener.NewFtpListener(env, 255)
	parsers := make([]reader.Parser, 0)
	parsers = append(parsers, &reader.SenseiveParser{})
	parsers = append(parsers, &reader.TrimbleParser{})

	for {
		filepath, err := ftpListener.Listen()
		if err != nil {
			log.Printf("Error listening for new files: %s\n", err.Error())
			panic(err)
		}
		log.Printf("Process file %s\n", filepath)
		parser, measures := process.FindParser(reader.ReadAndDelete(pkg.FtpLocalPath+"/"+filepath), parsers)
		if parser == nil {
			log.Printf("No parser found for file %s\n", filepath)
			ftpListener.RegisterBlacklist(filepath)
			continue
		}
		ftpListener.DeleteFile(filepath)
		go influxStorer.Store(parser.ExtractProject(filepath), parser.Source(), measures)
	}
}
