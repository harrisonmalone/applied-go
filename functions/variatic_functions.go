package main

import (
	"fmt"
	"unicode"
)

func longest(args ...string) int {
	return len(args)
}

func looping() {
	s := "abcde"
	for _, letter := range s {
			upperLetter := unicode.ToUpper(letter)
			fmt.Print(string(upperLetter))
	}
	fmt.Println("\n" + s)
}


func main() {
	// fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
	looping()
}