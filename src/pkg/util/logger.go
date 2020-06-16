package util

import (
	"fmt"
	"time"
)

func InfoLog(msg string, args ...interface{}) {
	currTime := time.Now().Format("01/09/1997 00:00:00")
	fmt.Printf("%v::INFO LOG  >> " + fmt.Sprintf(msg, args) + "\n", currTime)
}

func ErrorLog(msg string, err error, args ...interface{}) {
	currTime := time.Now().Format("01/09/1997 00:00:00")
	fmt.Printf("%v::ERROR LOG >> " + fmt.Sprintf(msg, args) + ":\n%v\n", currTime, err)
}

func DebugLog(msg string, args ...interface{}) {
	currTime := time.Now().Format("01/09/1997 00:00:00")
	fmt.Printf("%v::DEBUG LOG >> " + fmt.Sprintf(msg, args) + "\n", currTime)
}
