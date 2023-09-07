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
	localPath  string
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
	}
	return &FtpListener{
		client:     client,
		localPath:  env.FtpLocalPath,
		serverPath: env.FtpServerPath,
	}
}

func (f *FtpListener) Listen() (string, string) {
	log.Printf("Listening for new files... ")
	todo, filename := f.nextFile()
	for todo, filename = f.nextFile(); !todo; {
		if !todo {
			time.Sleep(pkg.WaitTime)
		}
	}
	log.Printf("File %s found\n", filename)
	f.downloadFile(filename)
	f.deleteFile(filename)
	return f.extractProject(filename), f.localPath + "/" + filename

}

func (f *FtpListener) nextFile() (bool, string) {
	files, err := f.client.ReadDir(f.serverPath)
	if err != nil {
		sentry.CaptureException(err)
	}
	if len(files) == 0 {
		return false, ""
	}
	return true, files[0].Name()
}

func (f *FtpListener) downloadFile(filename string) {
	localFile, err := os.Create(f.localPath + "/" + filename)
	if err != nil {
		sentry.CaptureException(err)
	}
	err = f.client.Retrieve(f.serverPath+"/"+filename, localFile)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func (f *FtpListener) deleteFile(filename string) {
	err := f.client.Delete(f.serverPath + "/" + filename)
	if err != nil {
		sentry.CaptureException(err)
	}
}

// Sample:  Geosud-Demo_rail_2023-09-06_14-05-53.csv
// project name: Geosud-Demo_rail
// split ath _20
func (f *FtpListener) extractProject(filename string) string {
	return strings.Split(filename, "_20")[0]
}
