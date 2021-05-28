package main

import (
	"fmt"

	defer_something "github.com/harrisonmalone/applied-go/defer"
)

// Add code to newClosure so that it returns two closures. The first one is of type func(), the second one of type func() int.

// Both closures shall modify an integer variable defined in the outer function as follows:

// The first closure shall just set the outer variable to 5. It returns nothing.

// The second closure shall multiply the outer variable by 7 and return the value.

// main() calls newClosure to create the new closures, and then calls both closures and prints out the result.

func newClosures() (func(), func() int) {
    a := 0
    closureOne := func () {
			a = 5
		}
		closureTwo := func() int {
			a = a * 7
			return a
		}
		return closureOne, closureTwo
}

func main() {
    f1, f2 := newClosures()
    f1() // a = 5
    n := f2() // multiplies "a" by 7 - is f2's internal value of "a" 0 or 5 before the call? it's 5
    fmt.Println(n)
		fmt.Println("Before f")
		defer_something.F()
		fmt.Println("After f")
}