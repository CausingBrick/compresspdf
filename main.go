package main

import (
	"container/list"
	"log"

	"github.com/astaxie/beego/config"
)

var (
	gs          string
	refreshTime string
)

func init() {
	cnf, err := config.NewConfig("ini", "settings.conf")
	if err != nil {
		log.Panicln(err)
	}
	gs = cnf.String("gsPath")
}

type pdfInfo struct {
}

type queue list.List

func (q *queue) Init() {
}

func main() {
	log.Println("Starting service.")
	for {
		compress(gs, "datastructer.pdf", "in.pdf", 4)
	}
}
