package main

import (
	"fmt"
	"strings"
)

// Implement a function that receives a string and a map[string]int. The function shall split the string into words, turn the words into lowercase, and adds the word to the map and/or increases the counter for this word.

func count(s string, m map[string]int) map[string]int {
	words := strings.Split(s, " ")
	for _, word := range words {
		lowerCaseWord := strings.ToLower(strings.Trim(word, "\t\n\"'.,:;?!()-"))
		m[lowerCaseWord]++
	}
	return m
}

// Write a function that receives a map[string]int and prints out each key (the word) together with its value (the count), but only if the count is greater than 1.

func printCount(m map[string]int) {
	for key, value := range m {
		if value > 1 {
			fmt.Println(key, value)
		}
	}
}

func main() {
	result := count("Harrison Malone Harrison Animal Nelson Bee", map[string]int{})
	printCount(result)
}