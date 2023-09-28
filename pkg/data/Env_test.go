package data

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	err := os.Setenv("INFLUX_TOKEN", "token")
	if err != nil {
		return
	}
	err = os.Setenv("INFLUX_URL", "url")
	if err != nil {
		return
	}
	err = os.Setenv("INFLUX_ORG", "org")
	if err != nil {
		return
	}
	err = os.Setenv("INFLUX_PREFIX", "prefix")
	if err != nil {
		return
	}
	err = os.Setenv("FTP_SERVER_URL", "ftpserver")
	if err != nil {
		return
	}
	err = os.Setenv("FTP_SERVER_PATH", "ftppath")
	if err != nil {
		return
	}
	err = os.Setenv("FTP_USER", "ftpuser")
	if err != nil {
		return
	}
	err = os.Setenv("FTP_PASSWORD", "ftppassword")
	if err != nil {
		return
	}

	env := ReadEnv()

	if env.InfluxToken != "token" {
		t.Errorf("env.InfluxToken = %s; want token", env.InfluxToken)
	}
	if env.InfluxUrl != "url" {
		t.Errorf("env.InfluxUrl = %s; want url", env.InfluxUrl)
	}
	if env.InfluxOrg != "org" {
		t.Errorf("env.InfluxOrg = %s; want org", env.InfluxOrg)
	}
	if env.InfluxPrefix != "prefix" {
		t.Errorf("env.InfluxPrefix = %s; want prefix", env.InfluxPrefix)
	}
	if env.FtpServer != "ftpserver" {
		t.Errorf("env.FtpServer = %s; want ftpserver", env.FtpServer)
	}
	if env.FtpServerPath != "ftppath" {
		t.Errorf("env.FtpServerPath = %s; want ftppath", env.FtpServerPath)
	}
	if env.FtpUser != "ftpuser" {
		t.Errorf("env.FtpUser = %s; want ftpuser", env.FtpUser)
	}
	if env.FtpPassword != "ftppassword" {
		t.Errorf("env.FtpPassword = %s; want ftppassword", env.FtpPassword)
	}

	expected := "InfluxToken: token\n" +
		"InfluxUrl: url\n" +
		"InfluxOrg: org\n" +
		"InfluxPrefix: prefix\n" +
		"FtpServer: ftpserver\n" +
		"FtpServerPath: ftppath\n" +
		"FtpUser: ftpuser\n" +
		"FtpPassword: ftppassword\n"
	if env.String() != expected {
		t.Errorf("env.String() = %s; want %s", env.String(), expected)
	}

	err = os.Unsetenv("INFLUX_TOKEN")
	if err != nil {
		return
	}
	err = os.Unsetenv("INFLUX_URL")
	if err != nil {
		return
	}
	err = os.Unsetenv("INFLUX_ORG")
	if err != nil {
		return
	}
	err = os.Unsetenv("INFLUX_PREFIX")
	if err != nil {
		return
	}
	err = os.Unsetenv("FTP_SERVER_URL")
	if err != nil {
		return
	}
	err = os.Unsetenv("FTP_SERVER_PATH")
	if err != nil {
		return
	}
	err = os.Unsetenv("FTP_USER")
	if err != nil {
		return
	}
	err = os.Unsetenv("FTP_PASSWORD")
	if err != nil {
		return
	}
}
