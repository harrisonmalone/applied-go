package main

import (
	"flag"
	"fmt"
)

func main() {
	
	// A floating-point variable that holds the value to convert.
	var input float64
	// fmt.Println(input)
	
	// // A floating-point value that holds the result.
	var result float64
	// fmt.Println(result)

	// Set up and parse the -from and -to flags.
	// The arguments to flag.StringVar() are: A pointer to the target
	// variable, the name, the default value, and a description.
	/* TODO: Your code here */
	var fromType, toType string
	flag.StringVar(&fromType, "from", "", "from type")
	flag.StringVar(&toType, "to", "", "to type")
	flag.Parse()
	
	// Let's look how the arguments look like...
	
	//...from the os package's perspective
	// fmt.Println("os.Args:", os.Args)
	
	//...and from the flag package's perspective (after having parsed the flags):
	// fmt.Println("flag.Args():", flag.Args())
	
	// Check the -from and -to flags. If empty, print an error message
	// and exit.
	/* TODO: Your code here */
	if fromType == "" {
		fmt.Println("no from type passed in as arg")
		return
	}

	if toType == "" {
		fmt.Println("no to type")
		return
	}

	// Check if an argument is passed. If not, print a
	// message and exit.
	/* TODO: Your code here */
	if len(flag.Args()) == 0 {
		fmt.Println("no arg passed")
		return
	}

	// Scan the input value (remember this should be the only remaining
	// value in flag.Args() after the flags were parsed).
	//
	// fmt.Sscanf is the variant of Scanf that scans from a string.
	// Its arguments are: The string to scan, the format string,
	// and one or more pointer to the variable(s) to fill.
	//
	// flag.Args() is a slice of strings. Use the index operator [n]
	// to fetch the string at position n.
	numScanned, err := fmt.Sscanf(flag.Args()[0], "%f", &input)
	fmt.Println(numScanned)
	if numScanned != 1 || err != nil {
		fmt.Println("Cannot scan the value to convert. flag.Args() is:", flag.Args())
		return
	}

	// Convert the value based on the from and to units.
	// Here, we use a switch statement to select the formula
	// depending on the from and to units.
	// Let's keep things simple and only convert from:
	//
	// * meters to feet,
	// * feet to meters,
	// * meters to inches, and
	// * inches to meters.
	//
	/* TODO: Your code here */
	switch {
		case fromType == "feet" && toType == "meters":
			result = input / 3.28
		case fromType == "meters" && toType == "feet":
			result = input * 3.28
		case fromType == "meters" && toType == "inches":
			result = input * 39.37
		case fromType == "inches" && toType == "meters":
			result = input / 39.37
	}

	// Finally, print the result.
	/* TODO: Your code here */
	fmt.Printf("The result is %.2f %s\n", result, toType)
}