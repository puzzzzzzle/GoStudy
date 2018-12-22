package main

import (
	"fmt"
	"os"
)

func IsTen(key int) (int, error) {
	if key != 10 {
		return -1, fmt.Errorf("value(%d) is not 10", key)
	}
	return 10 * 10, nil
}
func main() {
	var in int
	_, in_err := fmt.Fscan(os.Stdin, &in)
	if in_err != nil {
		panic("in_err!=nil")
	}
	out, err := IsTen(in)
	if err != nil {
		fmt.Println(err)
		panic("unlucky!!!")
	}
	fmt.Println(out)
}
