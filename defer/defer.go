package defer_something

import "fmt"

// Write a function trace() that receives a string - the name of the current function - and does the following:

// Print “Entering <name>” where <name> is the string parameter that trace receives
// Create and return a function that prints “Leaving <name>”
// Then call trace() via the defer keyword in such a way that trace() runs immediately, and returns its result to defer.

func trace(name string) func() {
	// TODO:
	// 1. Print "Entering <name>"
	fmt.Println(fmt.Sprintf("Entering %s", name))
	// 2. return a func() that prints "Leaving <name>" 
	return func() {
		fmt.Println(fmt.Sprintf("Leaving %s", name))
	}
}

func F() {
	defer trace("f")()
	fmt.Println("Doing something")
}