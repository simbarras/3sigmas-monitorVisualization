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
	log.Printf("Launching trimble reader with env: \n%+v\n", env)

	influxStorer := storer.NewInfluxStorer(env)
	ftpListener := listener.NewFtpListener(env, 255)
	parser := reader.TrimbleParser{}

	process.Process(ftpListener, influxStorer, &parser)

}
