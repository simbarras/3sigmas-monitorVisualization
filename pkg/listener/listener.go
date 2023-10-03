package listener

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/secsy/goftp"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
	"os"
	"strings"
	"time"
)

type FtpListener struct {
	client     *goftp.Client
	serverPath string
	blacklist  []string
	index      int
	size       int
}

func NewFtpListener(env data.Env, maxIndex int) *FtpListener {
	config := goftp.Config{
		User:     env.FtpUser,
		Password: env.FtpPassword,
	}
	client, err := goftp.DialConfig(config, env.FtpServer)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return &FtpListener{
		client:     client,
		serverPath: env.FtpServerPath,
		blacklist:  make([]string, maxIndex),
		index:      -1,
		size:       0,
	}
}

func (f *FtpListener) Listen() (string, error) {
	log.Printf("Listening for new files... ")
	todo, filename, err := f.nextFile()
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}
	for !todo {
		if !todo {
			time.Sleep(pkg.WaitTime)
		}
		todo, filename, err = f.nextFile()
		if err != nil {
			sentry.CaptureException(err)
			return "", err
		}
	}
	log.Printf("File %s found\n", filename)
	f.downloadFile(filename)
	return filename, nil

}

func (f *FtpListener) nextFile() (bool, string, error) {
	if f.client == nil {
		return false, "", errors.New("no connection established")
	}
	files, err := f.client.ReadDir(f.serverPath)
	if err != nil {
		sentry.CaptureException(err)
		return false, "", err
	}
	// Filter to keep only csv files
	files = filter(files, func(file os.FileInfo) bool {
		if !strings.HasSuffix(file.Name(), ".csv") {
			return false
		}
		for _, filename := range f.blacklist {
			if filename == file.Name() {
				return false
			}
		}
		return true
	})
	if err != nil {
		sentry.CaptureException(err)
		return false, "", err
	}
	if len(files) == 0 {
		return false, "", nil
	}
	log.Printf("Found %d file(s)\n", len(files))
	return true, files[0].Name(), nil
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

func (f *FtpListener) DeleteFile(filename string) {
	err := f.client.Delete(f.serverPath + "/" + filename)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	log.Printf("File %s deleted on FTP\n", filename)
}

func (f *FtpListener) RegisterBlacklist(filename string) {
	f.index = (f.index + 1) % len(f.blacklist)
	f.blacklist[f.index] = filename
	f.size++
	if f.size > len(f.blacklist) {
		f.size = len(f.blacklist)
	}
	log.Printf("File %s registered in blacklist at index %d with size %d and max size %d\n", filename, f.index, f.size, len(f.blacklist))
}

func filter[K comparable](files []K, filter func(K) bool) []K {
	var filtered []K
	for _, file := range files {
		if filter(file) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}
