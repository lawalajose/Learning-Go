package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func calculate(num int) int {
	return num + 1
}

func main() {
	switch result := calculate(5); {
	case result >= 10:
		fmt.Println(">10")
	case result == 6:
		fmt.Println("==6")
	case result <= 10:
		fmt.Println("<=10")
	}

}
