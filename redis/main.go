package main

import (
	"fmt"
	rd "github.com/garyburd/redigo/redis"
	"os"
)

func main() {
	conn, err := rd.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("redis conn errr with error : %s", err.Error())
		return
	}

	defer conn.Close()

	_, err = conn.Do("SET", "GOPID", os.Getpid())
	if err != nil {
		fmt.Println("set pid err" + err.Error())
		return
	} else {
		fmt.Println("set pid ok")
	}

	pid, err := rd.Int(conn.Do("GET", "GOPID"))
	if err != nil {
		fmt.Println("get pid wrong")
	} else {
		fmt.Println(pid)
	}

	_, err = conn.Do("EXPIRE", "GOPID", 10)
	if err != nil {
		fmt.Println("expire wrong :" + err.Error())
	}

}
