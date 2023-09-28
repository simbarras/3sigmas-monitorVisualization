package listener

import (
	"3sigmas-monitorVisualization/pkg/data"
	"strings"
	"testing"
)

func TestNewFtpListener(t *testing.T) {
	env := data.Env{
		FtpUser:     "user",
		FtpPassword: "password",
		FtpServer:   "server",
	}
	f := NewFtpListener(env, 3)
	if f != nil {
		t.Errorf("No connection should be established with dummy configuration")
	}
}

func TestFtpListener_Listen(t *testing.T) {
	f := FtpListener{
		client:     nil,
		serverPath: "serverPath",
		blacklist:  []string{"file1.csv", "file2.csv", "file3.csv"},
		index:      0,
		size:       3,
	}
	_, err := f.Listen()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFtpListener_RegisterBlacklist(t *testing.T) {
	const length = 3
	f := FtpListener{
		blacklist: make([]string, length),
		index:     -1,
		size:      0,
	}
	if f.index != -1 {
		t.Errorf("Expected -1, got %d", f.index)
	}
	if f.size != 0 {
		t.Errorf("Expected 0, got %d", f.size)
	}
	if len(f.blacklist) != length {
		t.Errorf("Expected %d, got %d", length, len(f.blacklist))
	}

	f.RegisterBlacklist("file1.csv")
	if f.index != 0 {
		t.Errorf("Expected 0, got %d", f.index)
	}
	if f.size != 1 {
		t.Errorf("Expected 1, got %d", f.size)
	}
	if f.blacklist[0] != "file1.csv" {
		t.Errorf("Expected file1.csv, got %s", f.blacklist[0])
	}

	f.RegisterBlacklist("file2.csv")
	if f.index != 1 {
		t.Errorf("Expected 1, got %d", f.index)
	}
	if f.size != 2 {
		t.Errorf("Expected 2, got %d", f.size)
	}
	if f.blacklist[1] != "file2.csv" {
		t.Errorf("Expected file2.csv, got %s", f.blacklist[1])
	}

	f.RegisterBlacklist("file3.csv")
	if f.index != 2 {
		t.Errorf("Expected 2, got %d", f.index)
	}
	if f.size != 3 {
		t.Errorf("Expected 3, got %d", f.size)
	}
	if f.blacklist[2] != "file3.csv" {
		t.Errorf("Expected file3.csv, got %s", f.blacklist[2])
	}

	f.RegisterBlacklist("file4.csv")
	if f.index != 0 {
		t.Errorf("Expected 0, got %d", f.index)
	}
	if f.size != 3 {
		t.Errorf("Expected 3, got %d", f.size)
	}
	if f.blacklist[0] != "file4.csv" {
		t.Errorf("Expected file4.csv, got %s", f.blacklist[0])
	}
}

func TestFtpListener_filter(t *testing.T) {
	files := []string{"file1.csv", "file2.csv", "file3.txt", "file4.csv"}
	filteredFiles := filter(files, func(file string) bool {
		if !strings.HasSuffix(file, ".csv") {
			return false
		}
		return true
	})
	if len(filteredFiles) != 3 {
		t.Errorf("Expected 3 files, got %d", len(filteredFiles))
	}
	if filteredFiles[0] != "file1.csv" {
		t.Errorf("Expected file1.csv, got %s", filteredFiles[0])
	}
	if filteredFiles[1] != "file2.csv" {
		t.Errorf("Expected file2.csv, got %s", filteredFiles[1])
	}
	if filteredFiles[2] != "file4.csv" {
		t.Errorf("Expected file4.csv, got %s", filteredFiles[2])
	}

	files2 := []int{10, 11, 12, 21}
	filteredFiles2 := filter(files2, func(file int) bool {
		if file%2 == 0 {
			return false
		}
		return true
	})
	if len(filteredFiles2) != 2 {
		t.Errorf("Expected 2 files, got %d", len(filteredFiles2))
	}
	if filteredFiles2[0] != 11 {
		t.Errorf("Expected 11, got %d", filteredFiles2[0])
	}
	if filteredFiles2[1] != 21 {
		t.Errorf("Expected 21, got %d", filteredFiles2[1])
	}
}
