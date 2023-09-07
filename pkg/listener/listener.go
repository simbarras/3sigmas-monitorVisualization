package listener

import (
	"3sigmas-monitorVisualization/pkg"
	"fmt"
	"github.com/secsy/goftp"
	"os"
	"strings"
	"time"
)

type FtpListener struct {
	client *goftp.Client
}

func NewFtpListener() *FtpListener {
	config := goftp.Config{
		User:     pkg.FtpUser,
		Password: pkg.FtpPassword,
	}
	client, err := goftp.DialConfig(config, pkg.FtpServerUrl)
	if err != nil {
		panic(err)
	}
	return &FtpListener{
		client: client,
	}
}

func (f *FtpListener) Listen() (string, string) {
	todo, filename := f.nextFile()
	for todo, filename = f.nextFile(); !todo; {
		if !todo {
			fmt.Println("No file to read")
			time.Sleep(5 * time.Second)
		}
	}
	fmt.Println("Reading file: ", filename)
	f.downloadFile(filename)
	fmt.Println("File downloaded")
	f.deleteFile(filename)
	fmt.Println("File deleted")
	return f.extractProject(filename), pkg.FtpLocalPath + "/" + filename

}

func (f *FtpListener) nextFile() (bool, string) {
	files, err := f.client.ReadDir(pkg.FtpServerPath)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		return false, ""
	}
	return true, files[0].Name()
}

func (f *FtpListener) downloadFile(filename string) {
	localFile, err := os.Create(pkg.FtpLocalPath + "/" + filename)
	if err != nil {
		panic(err)
	}
	err = f.client.Retrieve(pkg.FtpServerPath+"/"+filename, localFile)
	if err != nil {
		panic(err)
	}
}

func (f *FtpListener) deleteFile(filename string) {
	err := f.client.Delete(pkg.FtpServerPath + "/" + filename)
	if err != nil {
		panic(err)
	}
}

// Sample:  Geosud-Demo_rail_2023-09-06_14-05-53.csv
// project name: Geosud-Demo_rail
// split ath _20
func (f *FtpListener) extractProject(filename string) string {
	return strings.Split(filename, "_20")[0]
}
