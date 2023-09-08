package listener

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/data"
	"github.com/getsentry/sentry-go"
	"github.com/secsy/goftp"
	"log"
	"os"
	"strings"
	"time"
)

type FtpListener struct {
	client     *goftp.Client
	serverPath string
}

func NewFtpListener(env data.Env) *FtpListener {
	config := goftp.Config{
		User:     env.FtpUser,
		Password: env.FtpPassword,
	}
	client, err := goftp.DialConfig(config, env.FtpServer)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	return &FtpListener{
		client:     client,
		serverPath: env.FtpServerPath,
	}
}

func (f *FtpListener) Listen() (string, string) {
	log.Printf("Listening for new files... ")
	todo, filename := f.nextFile()
	for !todo {
		if !todo {
			time.Sleep(pkg.WaitTime)
		}
		todo, filename = f.nextFile()
	}
	log.Printf("File %s found\n", filename)
	f.downloadFile(filename)
	f.deleteFile(filename)
	return f.extractProject(filename), pkg.FtpLocalPath + "/" + filename

}

func (f *FtpListener) nextFile() (bool, string) {
	files, err := f.client.ReadDir(f.serverPath)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	if len(files) == 0 {
		return false, ""
	}
	log.Printf("Found %d file(s)\n", len(files))
	return true, files[0].Name()
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}

func (f *FtpListener) downloadFile(filename string) {
	localFile, err := os.Create(pkg.FtpLocalPath + "/" + filename)
	defer closeFile(localFile)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	err = f.client.Retrieve(f.serverPath+"/"+filename, localFile)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}

func (f *FtpListener) deleteFile(filename string) {
	err := f.client.Delete(f.serverPath + "/" + filename)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
}

// Sample:  Geosud-Demo_rail_2023-09-06_14-05-53.csv
// project name: Geosud-Demo_rail
// split at _20
func (f *FtpListener) extractProject(filename string) string {
	return strings.Split(filename, "_20")[0]
}
