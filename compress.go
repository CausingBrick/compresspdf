package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

/*
/Compression levels:
0: default
1: prepress
2: printer
3: ebook
4: screen*/
var quality = []string{"/default", "/prepress", "/printer", "/ebook", "/screen"}

func compress(gs, inFilePath, outFilePath string, power int) {
	arg := []string{"-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=" + quality[power],
		"-dNOPAUSE", "-dQUIET", "-dBATCH",
		"-sOutputFile=" + outFilePath,
		inFilePath}

	cmd := exec.Command(gs, arg...)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	fmt.Println(cmd.String())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: Hello World!
}
