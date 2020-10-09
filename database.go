package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// Defined pdf compressed state
const (
	UnCompressed = iota
	Compressed
	ErrorCompress
)

// GetPdfInfos returns the pdf info from database.
func GetPdfInfos() (*[]orm.Params, error) {
	o := orm.NewOrm()
	queryStr := "Select " + appConf.PDFPK + "," + appConf.PDFInput +
		" from " + appConf.TBName +
		" where " + appConf.PDFCompressState + " ='" + strconv.Itoa(UnCompressed) + "'"
	fmt.Println(UnCompressed, queryStr)

	//Query datas.
	res := new([]orm.Params)
	if _, err := o.Raw(queryStr).Values(res); err != nil {
		return res, errors.New("No pdf info found: " + err.Error())
	}
	return res, nil
}

func getDatasource() string {
	return appConf.DBUser + ":" + appConf.DBPsd + "@tcp(" + appConf.DBIP + ":" + appConf.DBPort + ")/" + appConf.DBName + "?charset=utf8"
}

// UpdatePdfState updates the state of pdf info.
func UpdatePdfState(guid string, state int, desPath string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE " + appConf.TBName + " SET state ='" + strconv.Itoa(state) + "' , compressed_path ='" + desPath + "'Where guid ='" + guid + "'").Exec()
	return err
}
