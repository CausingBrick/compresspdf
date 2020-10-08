package main

import (
	"log"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const configFile = "settings.conf"

var (
	//PdfInfoCh Mission queue
	PdfInfoCh chan orm.Params
	appConf   *Conf
)

func init() {
	var err error
	appConf, err = GetSettings(configFile)
	if err != nil {
		log.Panicln("Get config file error:\n", err)
	}
	PdfInfoCh = make(chan orm.Params, appConf.HandelNum)
}

func main() {
	log.Println("Starting service.")
	for {
		log.Println("Retireing datas...")
		pdfInfos, err := GetPdfInfos()
		CheckErr(err)
		//handel compersspdf if mission existed.
		if len(*pdfInfos) != 0 {
			log.Printf("Got %d rows, start to compressing", len(*pdfInfos))

			go func() {
				for _, v := range *pdfInfos {
					PdfInfoCh <- v
				}
			}()
			var wg sync.WaitGroup
			for range PdfInfoCh {
				go Compress(<-PdfInfoCh, &wg)
				wg.Add(1)
			}
			wg.Wait()
		} else { //wait for refreash if no mission.
			log.Println("No mission now, sleep for next check.")
			time.Sleep(time.Minute)
		}
	}
}
