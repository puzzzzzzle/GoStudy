package main

import (
	"fmt"
)

func main() {
	var i_arr [10]int64
	for k := range i_arr {
		i_arr[k] = int64(k)
	}
	fmt.Println(i_arr)

	var slice1 = make([]int64, 5)

	silce2 := make([]int64, 5)

	var slice3 []int64

	slice1 = []int64{0, 1, 2}

	copy(silce2, slice1[0:1])

	slice3 = append(slice3, 5, 6, 7, 8, 9)

	slice4 := i_arr[0:]

	slice5 := slice3[:]

	//remove
	slice5 = append(slice5[:2], slice5[2+1:]...)
	fmt.Println(slice5, slice4)

	students := make(map[string]float64)

	students["kyte"] = 82.5
	students["kar"] = 95.6
	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("%s->%f\n", k, v)
	}

	delete(students, "kar")
	delete(students, "kaar")
	fmt.Println(students["kar"])
	fmt.Println(students["kaar"])
	//fmt.Println(reflect.Type(students["kar"]))

}
