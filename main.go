package main

import (
	"log"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const configFile = "conf/settings.conf"

var (
	appConf *Conf
)

func init() {
	var err error
	appConf, err = GetSettings(configFile)
	if err != nil {
		log.Panicln(ErrorStr, "Get config file error:\n", err)
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", getDatasource())
}

func main() {
	log.Println(InfoStr, "Starting service.")
	for {
		log.Println(InfoStr, "Retireing datas...")

		//Get pdf info from database
		pdfInfos, err := GetPdfInfos()
		CheckErr(err)

		//handel compersspdf if mission existed.
		if dataLen := len(*pdfInfos); dataLen != 0 {
			log.Printf(InfoStr, "Got %d rows, start to compressing", dataLen)

			//Enqueue mission
			pdfInfoCh := make(chan orm.Params, dataLen)
			go func() {
				for _, v := range *pdfInfos {
					pdfInfoCh <- v
				}
				close(pdfInfoCh)
			}()

			var wg sync.WaitGroup
			// Set compression thread num
			wg.Add(appConf.HandelNum)

			//Start compression accroding to the HandelNum
			for i := 0; i < appConf.HandelNum; i++ {
				go Compress(pdfInfoCh, &wg)
			}
			//wait mission done
			wg.Wait()
			log.Println(InfoStr, "Mission finshed!")

		} else { //wait for refreash if no mission.
			log.Println(InfoStr, "No mission now, sleep for", int(appConf.RefreshMinute), "minuties.")
			time.Sleep(appConf.RefreshMinute * time.Minute)
		}
	}
}
