package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/listener"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/process"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/reader"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/storer"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/trigger"
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
		go func() {
			bucketName, err := influxStorer.Store(parser.ExtractProject(filepath), parser.Source(), measures)
			if err != nil {
				log.Printf("Error storing measures: %s\n", err.Error())
				sentry.CaptureException(err)
			} else {
				err := trigger.Trigger(env.ApiTrigger, bucketName)
				if err != nil {
					log.Printf("Error triggering: %s\n", err.Error())
					sentry.CaptureException(err)
				}
			}
		}()
	}
}
