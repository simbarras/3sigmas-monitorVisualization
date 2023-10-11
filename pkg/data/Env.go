package data

import "os"

type Env struct {
	InfluxToken   string
	InfluxUrl     string
	InfluxOrg     string
	InfluxPrefix  string
	FtpServer     string
	FtpServerPath string
	FtpUser       string
	FtpPassword   string
	ApiTrigger    string
}

func (e *Env) String() string {
	return "InfluxToken: " + e.InfluxToken + "\n" +
		"InfluxUrl: " + e.InfluxUrl + "\n" +
		"InfluxOrg: " + e.InfluxOrg + "\n" +
		"InfluxPrefix: " + e.InfluxPrefix + "\n" +
		"FtpServer: " + e.FtpServer + "\n" +
		"FtpServerPath: " + e.FtpServerPath + "\n" +
		"FtpUser: " + e.FtpUser + "\n" +
		"FtpPassword: " + e.FtpPassword + "\n"
}

func ReadEnv() Env {
	env := Env{
		InfluxToken:   os.Getenv("INFLUX_TOKEN"),
		InfluxUrl:     os.Getenv("INFLUX_URL"),
		InfluxOrg:     os.Getenv("INFLUX_ORG"),
		InfluxPrefix:  os.Getenv("INFLUX_PREFIX"),
		FtpServer:     os.Getenv("FTP_SERVER_URL"),
		FtpServerPath: os.Getenv("FTP_SERVER_PATH"),
		FtpUser:       os.Getenv("FTP_USER"),
		FtpPassword:   os.Getenv("FTP_PASSWORD"),
		ApiTrigger:    os.Getenv("API_TRIGGER"),
	}
	return env
}
