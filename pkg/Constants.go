package pkg

import (
	"os"
	"time"
)

const (
	Version        = "0.1.2"
	WaitTime       = 5 * time.Second
	FtpLocalPath   = "."
	SenseiveSource = "-senseive"
)

func Filter(files []os.FileInfo, filter func(os.FileInfo) bool) []os.FileInfo {
	var filtered []os.FileInfo
	for _, file := range files {
		if filter(file) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}
