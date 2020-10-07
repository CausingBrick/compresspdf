package main

import (
	"container/list"
	"log"
)

// CheckErr checks if err is NULL
func CheckErr(err error, info ...string) {
	if err != nil {
		log.Println(info, err)
	}
}

// Queue implements
type Queue struct {
	list.List
}

// EnQueue EnQueues
func (q *Queue) EnQueue(v interface{}) {
	q.PushBack(v)
}

// DeQueue DeQueues
func (q *Queue) DeQueue(pdfinfo *PdfInfo) interface{} {
	return (q.Remove(q.Back()))
}
