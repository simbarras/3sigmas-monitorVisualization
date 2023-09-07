package main

import (
	"3sigmas-monitorVisualization/pkg/listener"
	"3sigmas-monitorVisualization/pkg/reader"
	"3sigmas-monitorVisualization/pkg/storer"
	"fmt"
)

func main() {
	influxStorer := storer.NewInfluxStorer()
	ftpListener := listener.NewFtpListener()

	for {
		project, filepath := ftpListener.Listen()
		fmt.Printf("Process file %s from project %s\n", filepath, project)
		go influxStorer.Store(project, reader.SenseiveParse(reader.Read(filepath)))
	}
}
