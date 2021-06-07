package main

import "fmt"

func main() {
	src := []int{}
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	fmt.Printf("%p %[1]v\n", src)
	dest1 := append(src, 3)
	fmt.Printf("%p %[1]v\n", dest1)
	dest2 := append(src, 4)
	fmt.Printf("%p %[1]v\n", dest2)
	fmt.Println(src, dest1, dest2)
}