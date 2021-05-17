package main
import (
    "fmt"
    "os"
    "strconv"
)
func collatz(n int) int {
    count := 0
    for n > 1 {
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
        count++
    }
    return count
}

func main() {
    var n int
    var err error
    if len(os.Args) > 1 {
        n, err = strconv.Atoi(os.Args[1])
        if err != nil {
            fmt.Println(err)
            return
        }
    } else {
        fmt.Println("Input a number > 1: ")
        _, err := fmt.Scanf("%d", &n)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    if n <= 1 {
        fmt.Println("n must be larger than 1.")
        return
    }
    c := collatz(n)
    fmt.Println(n, "requires", c, "steps to reach 1.")
}