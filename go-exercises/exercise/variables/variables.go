//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

// ✅ TASK 1 — Parse an integer from user input

// Write a function:

// func ReadInt(input string) (int, error)

// Requirements:

// It takes a string (e.g., "42").

// It converts it to an integer using strconv.Atoi.

// If parsing fails, return a descriptive error like:

// invalid integer "abc": strconv.Atoi: parsing "abc": invalid syntax

// No printing inside the function — just return values.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInt(word string) (int, error) {
	num, err := strconv.Atoi(word)
	if err != nil {
		return 0, fmt.Errorf("invalid integer %q: %w", word, err)
	}
	return num, nil
}

func ReadFileName(path string) ([]byte, error) {
	word, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q %w", word, err)
	}
	return word, nil
}

//⭐ TASK 3 — Convert multiple strings to integers safely

// Write a function:

// func ParseInts(words []string) ([]int, error)

// Requirements:

// It takes a slice of strings, e.g., ["10", "20", "abc"].

// Convert each string to an integer using your ReadInt function.

// If any string fails:

// Return an error immediately (do not continue).

// Wrap the error with a message that includes the index of the failure.

// If all succeed, return the slice of ints and nil.

func ParseInts(words []string) ([]int, error) {
	var numbers []int
	for index, ch := range words {
		word, err := ReadInt(ch)
		numbers = append(numbers, word)
		if err != nil {
			return numbers, fmt.Errorf("Index of failure %q :%w", index, err)
		}
	}
	return numbers, nil
}

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by 0", a)
	}
	return a / b, nil
}

func main() {
	new, err := ReadFileName("kkls")
	fmt.Println(new, err)

}

func main() {

	words := strings.NewReader("Go is awesome")
	r := make([]byte, 4)
	word, _ err := r.Read(words)
	
}
