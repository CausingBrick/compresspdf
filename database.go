package main

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// GetPdfInfos returns the pdf info from database.
func GetPdfInfos() (*[]orm.Params, error) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", getDatasource())
	o := orm.NewOrm()
	queryStr := "Select * from " + appConf.TBName + " where compress_state ='0'"
	//Query datas.
	res := new([]orm.Params)
	if _, err := o.Raw(queryStr).Values(res); err != nil {
		return res, errors.New("No pdf info found: " + err.Error())
	}
	return res, nil
}

func getDatasource() string {
	fmt.Println(appConf)
	return appConf.DBUser + ":" + appConf.DBPsd + "@tcp(" + appConf.DBIP + ":" + appConf.DBPort + ")/" + appConf.DBName + "?charset=utf8"
}
