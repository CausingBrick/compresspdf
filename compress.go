package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"sync"

	"github.com/astaxie/beego/orm"
)

/*
/Compression levels:
0: default
1: prepress
2: printer
3: ebook
4: screen*/
var quality = []string{"/default", "/prepress", "/printer", "/ebook", "/screen"}

func compress(gs, inFilePath, outFilePath string, power int) error {
	arg := []string{"-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=" + quality[power],
		"-dNOPAUSE", "-dQUIET", "-dBATCH",
		"-sOutputFile=" + outFilePath,
		inFilePath}

	cmd := exec.Command(gs, arg...)
	fmt.Println(cmd.String())
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return err
}

// Compress handels pdf compression
func Compress(pdfInfoCh <-chan orm.Params, wait *sync.WaitGroup) {
	// Exception handling
	defer func() {
		if err := recover(); err != nil {
			log.Println(ErrorStr, "Catched Exception", err)
			return
		}
	}()
	for pdfInfo := range pdfInfoCh {
		fmt.Println(pdfInfo)
		//Creat outputFile name
		outputFile := path.Join(appConf.OutputPath, pdfInfo["guid"].(string)+".pdf")
		//start compress
		fmt.Println(appConf.GSPath, pdfInfo["file_path"].(string), outputFile)
		err := compress(appConf.GSPath, pdfInfo["file_path"].(string), outputFile, appConf.CompressLevel)
		if err != nil {
			log.Println(ErrorStr, "Error compression:", pdfInfo, err)
			UpdatePdfState(pdfInfo["guid"].(string), ErrorCompress, "")
			break
		}
		//Update pdf state that succeed to compress
		if err := UpdatePdfState(pdfInfo["guid"].(string), Compressed, outputFile); err != nil {
			log.Println(ErrorStr, "Error Update PdfInfo State:", err)
		}
	}
	wait.Done()
}
