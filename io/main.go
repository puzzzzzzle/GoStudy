package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "io_", log.Lshortfile|log.Ldate|log.Ltime)

func main() {
	file, err := os.Open("json.json")
	if err != nil {
		return
	}

	buff := bufio.NewReader(file)

	//byte,err := buff.ReadByte()

	//byte ,err:= ioutil.ReadFile(file.Name())

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(buff.ReadString('\n'))
	fmt.Println(file.Name())
}
