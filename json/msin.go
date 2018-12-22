package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type AddTestCel struct {
	LValue int `json:"lValue"`
	RValue int `json:"rValue"`
	Want   int `json:"want"`
}

var testValues = []AddTestCel{
	{5, 5, 10},
	{12, 15, 27},
	{4, 8, 12},
}

func main() {
	var jstr string
	for _, v := range testValues {
		jbytes, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		jstr += fmt.Sprintln(string(jbytes))
	}
	lines := strings.Split(jstr, "\n")
	for _, line := range lines {
		fmt.Println(line)
		var tmp AddTestCel
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println("json decode err : "+err.Error())
		}
		fmt.Println(tmp)
	}
	write,err :=os.OpenFile("json.json",os.O_CREATE|os.O_TRUNC|os.O_RDWR,0666)
	if err != nil {
		fmt.Println("create write err : "+err.Error())
	}
	defer write.Close()
	//defer write.Sync()

	write.WriteString(jstr)
	fmt.Fprintln(write,jstr)
}
