package main

import (
	"fmt"
	"os"
	"time"
)

func ping(seconds time.Duration, c chan<- string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * seconds)
		c <- fmt.Sprintf("ping on channel %d", seconds)
	}
}

func pong(stop chan bool) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping(1, channel1)
	go ping(2, channel2)

	for {
		select {
		case msg1 := <-channel1:
			fmt.Printf("received : %s\n", msg1)
		case msg2 := <-channel2:
			fmt.Printf("received : %s\n", msg2)
			//超时
		case <-time.After(3 * time.Second):
			fmt.Println("out time !!!")
			//停机
		case <-stop:
			fmt.Println("stop !!!")
			return
		}
	}
}

func main() {
	stop := make(chan bool)
	go pong(stop)
	var i int
	for {
		fmt.Println("input 1 to stop!")
		fmt.Fscan(os.Stdin, &i)
		if i == 1 {
			stop <- true
			break
		}
	}
	fmt.Println("bye")
}
