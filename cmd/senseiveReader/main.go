package main

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/data"
	"3sigmas-monitorVisualization/pkg/listener"
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
	log.Println("App started in release ", pkg.Version)

	env := data.ReadEnv()
	log.Printf("Launching with env: \n%+v\n", env)

	influxStorer := storer.NewInfluxStorer(env)
	ftpListener := listener.NewFtpListener(env)

	for {
		project, filepath := ftpListener.Listen()
		log.Printf("Process file %s from %s\n", filepath, project)
		go influxStorer.Store(project, reader.SenseiveParse(reader.Read(filepath)))
	}
}
