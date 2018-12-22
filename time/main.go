package main

import (
	"fmt"
	"time"
)

func main(){
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-09-13 23:38:02", local)

	t1 :=time.Now()

	fmt.Println(t,t1)
}
