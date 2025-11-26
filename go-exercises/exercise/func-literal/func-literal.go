//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import "fmt"

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	mul := func() {
		all := 0
		numlet := 0
		numdigit := 0
		numspaces := 0
		numpunc := 0
		for _, r := range lines {
			for _, i := range r {
				all += 1
				if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
					numlet += 1
				} else if i == ' ' {
					numspaces += 1
				} else if i >= '0' && i <= '9' {
					numdigit += 1
				} else {
					numpunc += 1
				}
			}
		}
		fmt.Printf(" Number of letters are : %v\n Number of digits are : %v\n Number of spaces are : %v\n Number of punctations are : %v\n All the runes are %v\n", numlet, numdigit, numspaces, numpunc, all)
	}
	mul()
}
