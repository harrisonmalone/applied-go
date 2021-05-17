package main

import (
	"flag"
	"fmt"
)

func main() {
	var max int
	flag.IntVar(&max, "max", 10, "Maximal sum")
	flag.Parse()
	fmt.Println(max)
}