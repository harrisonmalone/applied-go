package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func acronym(s string) (acr string) {
	afterSpace := false
	for i, letter := range s {
		if ((i == 0 || afterSpace) && unicode.IsLetter(letter)) {
			acr += string(letter)
			afterSpace = false
		}
		if unicode.IsSpace(letter) {
			afterSpace = true
		}
	}
	return acr
}

func main() {
    s := "Pan Galactic Gargle Blaster" 
    if len(os.Args) > 1 {
        s = strings.Join(os.Args, " ")
    }
    fmt.Println(acronym(s))
}