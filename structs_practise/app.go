package main

import (
	"bufio"
	"fmt"
	"os"
)

type stringSet map[string]struct{}
var exists = struct{}{}

func iter(n int) []struct{} {
	return make([]struct{}, n)
}

func createSet() stringSet {
	return make(map[string]struct{})
}

func add(s stringSet, value string) {
	s[value] = exists
}

func remove(s stringSet, value string) {
	delete(s, value)
}

func contains(s stringSet, value string) bool {
	_, c := s[value]
	return c
}

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("File cannot be opened", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	set := createSet()
	for scanner.Scan() {
		line := scanner.Text()
		if contains(set, line) {
			fmt.Println(line)	
		} else {
			add(set, line)
		}
	}
}
