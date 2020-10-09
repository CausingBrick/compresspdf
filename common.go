package main

import (
	"fmt"
	"log"
)

// Define print string kind
var (
	InfoStr  = Green("INFO  ")
	ErrorStr = Green("Error  ")
)

// CheckErr checks if err is NULL
func CheckErr(err error, info ...string) {
	if err != nil {
		log.Println(info, err)
	}
}

// Red returns a red string
func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

// Green returns a green string
func Green(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
}
