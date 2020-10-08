package main

import (
	"log"
)

// CheckErr checks if err is NULL
func CheckErr(err error, info ...string) {
	if err != nil {
		log.Println(info, err)
	}
}
