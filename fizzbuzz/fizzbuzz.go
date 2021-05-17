package main

import (
    "fmt"
    "os"
    "strconv"
)

func fizzbuzz(n int) {
    counter := 1
    for counter <= n {
        switch true {
        case (counter % 15 == 0):
            fmt.Println("FizzBuzz")
        case (counter % 3 == 0): 
            fmt.Println("Fizz")
        case (counter % 5 == 0): 
            fmt.Println("Buzz")
        default: 
            fmt.Println(counter)
        }
        counter++
    }
}

func main() {
    n := 50
    if len(os.Args) > 1 {
        max, err := strconv.Atoi(os.Args[1])
        if err == nil {
            n = max
        }
    }
    fizzbuzz(n)
}