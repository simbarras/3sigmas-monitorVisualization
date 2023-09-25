package main

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/data"
	"3sigmas-monitorVisualization/pkg/listener"
	"3sigmas-monitorVisualization/pkg/process"
	"3sigmas-monitorVisualization/pkg/reader"
	"3sigmas-monitorVisualization/pkg/storer"
	"log"
)

func main() {
	process.SetSentry()

	log.Printf("App started in release %s\n", pkg.Version)

	env := data.ReadEnv()
	log.Printf("Launching senseive reader with env: \n%+v\n", env)

	influxStorer := storer.NewInfluxStorer(env)
	ftpListener := listener.NewFtpListener(env, 255)
	parsers := make([]reader.Parser, 0)
	parsers = append(parsers, &reader.SenseiveParser{})
	parsers = append(parsers, &reader.TrimbleParser{})

	for {
		filepath := ftpListener.Listen()
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
