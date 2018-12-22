package main

import (
	"fmt"
)

func main() {
	var s1 = "hello"

	fmt.Println(s1)

	s2 := "s2"

	fmt.Printf("s2:%s\n", s2)

	var i1 int64 = 10

	fmt.Println(i1)

	var ai1 [10]int64

	ai1[0] = 10

	fmt.Println(&i1)

	var pi *int64 = &i1

	fmt.Println(pi)
	//fmt.Printf("count : %ld, sum :%ld",sumNumbers(12,5,5,6))
	fmt.Println(sumNumbers(12, 5, 8, 6))

	for i := 0; i < len(ai1); i++ {
		fmt.Println(ai1[i])
	}
	index := 0
	for index < 10 {
		fmt.Println(ai1[index])
		index++
	}
	for k, v := range ai1 {
		fmt.Print("key ", k)
		fmt.Println(" value ", v)
	}
	for _, v := range ai1 {
		//fmt.Print("key ", _)
		fmt.Println(" value ", v)
	}

	switch 3 {
	case 1, 2, 3, 4, 10:
		fmt.Printf("1,2,3,4,10\n")
	case 0:
		fmt.Printf("%d\n", 0)

	}
	switch 3 {
	case 1, 2, 3, 4, 10:
		fmt.Printf("1,2,3,4,10\n")
		fallthrough
	case 0:
		fmt.Printf("%d\n", 0)
	}
}

func sumNumbers(numbers ...int64) (count, sum int64) {
	defer fmt.Println("end")
	defer fmt.Println("end1")

	count = int64(len(numbers))
	sum = 0
	for _, e := range numbers {
		sum += e
	}
	return
}
