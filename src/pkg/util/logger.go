package util

import (
	"fmt"
	"time"
)

func InfoLog(msg string) {
	currTime := time.Now().Format("01/09/1997 00:00:00")
	fmt.Printf("%v::INFO LOG  >> " + msg + "\n", currTime)
}

func ErrorLog(msg string, err error) {
	currTime := time.Now().Format("01/09/1997 00:00:00")
	fmt.Printf("%v::ERROR LOG >> " + msg + ":\n%v\n", currTime, err)
}

func DebugLog(msg string) {
	currTime := time.Now().Format("01/09/1997 00:00:00")
	fmt.Printf("%v::DEBUG LOG >> " + msg + "\n", currTime)
}
