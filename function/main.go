package main

import (
	"errors"
	"fmt"
)

type mFloat float64



func hello(s string) (string, error) {
	if s == "" {
		return "", errors.New("none input")
	}
	return "hello " + s, nil
}

func main() {
	f := hello
	var f1 func(string) (string, error)
	f1 = f

	s, e := f1("")
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(s)

	f2 := func(s string) (string, error) {
		if s == "" {
			return "", errors.New("none input")
		}
		return "hello " + s, nil
	}

	f2("")
}
