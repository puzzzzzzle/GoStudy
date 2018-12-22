package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("%d : %s\n", i, v)
	}
	s := flag.String("s", "hello s", "input string!!!")
	i := flag.Int("i", 0, "int value")

	//clone :=flag.NewFlagSet("clone",flag.ExitOnError)

	flag.Parse()
	fmt.Printf("s : %s\n", *s)
	fmt.Printf("i : %d\n", *i)
}
