package util

import (
	"fmt"
	"time"
)

const (
	timeFormat = "01/02/2006 15:04:05"
)

func InfoLog(msg string) {
	currTime := time.Now().Format(timeFormat)
	fmt.Printf("%v::INFO LOG  >> " + msg + "\n", currTime)
}

func ErrorLog(msg string, err error) {
	currTime := time.Now().Format(timeFormat)
	fmt.Printf("%v::ERROR LOG >> " + msg + ":\n%v\n", currTime, err)
}

func DebugLog(msg string) {
	currTime := time.Now().Format(timeFormat)
	fmt.Printf("%v::DEBUG LOG >> " + msg + "\n", currTime)
}
