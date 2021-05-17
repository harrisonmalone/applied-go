package main

import (
	"os"
	"fmt"
)

func main() {
	for index, params := range os.Args {
		fmt.Println(index)
		fmt.Println(params)
	}
}