package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile ,err:=os.Create("log.log")
	if err != nil {
		fmt.Println("log file init err")
		return
	}
	defer logFile.Close()
	logger:=log.New(logFile,"logStudy",log.Ldate|log.Lmicroseconds|log.Llongfile)
	logger.SetOutput(os.Stdout)
	logger.Printf("tao log one %d\n",10)
}
