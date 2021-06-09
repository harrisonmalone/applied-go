package main

import "fmt"

type Counter int

func (c *Counter) add() {
	*c += 1
}

func (c *Counter) subtract() {
	*c -= 1
}

func main() {
	var counter Counter = 0
	fmt.Println(counter)
	counter.add()
	fmt.Println(counter)
	counter.subtract()
	fmt.Println(counter)
	counter.add()
	counter.add()
	counter.add()
	fmt.Println(counter)
}
