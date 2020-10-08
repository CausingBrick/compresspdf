package main

import (
	"time"

	"github.com/astaxie/beego/config"
)

// Conf maps the fields of settings.conf
type Conf struct {
	//GSPath is the path of gswin64.exe
	GSPath string
	//HandelNum is the pdf compression goroutine numbers.
	HandelNum     int
	InputPath     string
	OutputPath    string
	RefreshMinute time.Duration
	DBIP          string
	DBPort        string
	DBUser        string
	DBPsd         string
}

//GetSettings retires the datas in configFile
func GetSettings(configFile string) (*Conf, error) {
	cnf, err := config.NewConfig("ini", configFile)
	if err != nil {
		return nil, err
	}
	conf := &Conf{
		GSPath: cnf.String("gsPath"),
		DBIP:   cnf.String("dbIP"),
		DBPort: cnf.String("dbPort"),
		DBUser: cnf.String("dbUser"),
		DBPsd:  cnf.String("dbPsd"),
	}
	conf.HandelNum, err = cnf.Int("handelNum")
	if err != nil {
		return conf, err
	}
	refreshtime, err := cnf.Int("refreshMinute")
	if err != nil {
		return conf, err
	}
	conf.RefreshMinute = time.Duration(refreshtime)
	return conf, err
}
