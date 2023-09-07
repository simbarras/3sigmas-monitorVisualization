package data

import "os"

type Env struct {
	InfluxToken   string
	InfluxUrl     string
	InfluxOrg     string
	FtpServer     string
	FtpServerPath string
	FtpUser       string
	FtpPassword   string
	FtpLocalPath  string
}

func (e *Env) String() string {
	return "InfluxToken: " + e.InfluxToken + "\n" +
		"InfluxUrl: " + e.InfluxUrl + "\n" +
		"InfluxOrg: " + e.InfluxOrg + "\n" +
		"FtpServer: " + e.FtpServer + "\n" +
		"FtpServerPath: " + e.FtpServerPath + "\n" +
		"FtpUser: " + e.FtpUser + "\n" +
		"FtpPassword: " + e.FtpPassword + "\n" +
		"FtpLocalPath: " + e.FtpLocalPath + "\n"
}

func ReadEnv() Env {
	env := Env{
		InfluxToken:   os.Getenv("INFLUX_TOKEN"),
		InfluxUrl:     os.Getenv("INFLUX_URL"),
		InfluxOrg:     os.Getenv("INFLUX_ORG"),
		FtpServer:     os.Getenv("FTP_SERVER_URL"),
		FtpServerPath: os.Getenv("FTP_SERVER_PATH"),
		FtpUser:       os.Getenv("FTP_USER"),
		FtpPassword:   os.Getenv("FTP_PASSWORD"),
		FtpLocalPath:  os.Getenv("FTP_LOCAL_PATH"),
	}
	return env
}
