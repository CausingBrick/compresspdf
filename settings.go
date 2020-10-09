package main

import (
	"time"

	"github.com/astaxie/beego/config"
)

// Conf maps the fields of settings.conf
type Conf struct {
	//GSPath is the path of gswin64.exe
	GSPath     string
	InputPath  string
	OutputPath string
	//HandelNum is the pdf compression goroutine numbers.
	HandelNum int
	// 	# Compression levels:
	// #   0: default
	// #   1: prepress
	// #   2: printer
	// #   3: ebook
	// #   4: screen
	CompressLevel int
	// Time to retire pdf info data form database
	RefreshMinute time.Duration
	// Database info
	DBIP   string
	DBPort string
	DBUser string
	DBPsd  string
	DBName string
	TBName string
	//  These fileds map to table's filed
	PDFPK            string
	PDFInput         string
	PDFOutput        string
	PDFCompressState string
}

//GetSettings retires the datas in configFile
func GetSettings(configFile string) (*Conf, error) {
	cnf, err := config.NewConfig("ini", configFile)
	if err != nil {
		return nil, err
	}
	conf := &Conf{
		GSPath:           cnf.String("gsPath"),
		InputPath:        cnf.String("inputPath"),
		OutputPath:       cnf.String("outputPath"),
		DBIP:             cnf.String("dbIP"),
		DBPort:           cnf.String("dbPort"),
		DBUser:           cnf.String("dbUser"),
		DBPsd:            cnf.String("dbPsd"),
		DBName:           cnf.String("dbName"),
		TBName:           cnf.String("tbName"),
		PDFPK:            cnf.String("PDFPK"),
		PDFInput:         cnf.String("PDFInput"),
		PDFOutput:        cnf.String("PDFOutput"),
		PDFCompressState: cnf.String("pdfCompressState"),
	}
	conf.HandelNum, err = cnf.Int("handelNum")
	if err != nil {
		return conf, err
	}
	conf.CompressLevel, err = cnf.Int("compressLevel")
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
