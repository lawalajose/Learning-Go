//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operand int

const (
	Add Operand = iota
	Substract
	Multiply
	Divide
)

func (op Operand) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Substract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("Unaddressed problem")
}

func main() {
	add := Operand(Add)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Operand(Substract)
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Operand(Multiply)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Operand(Divide)
	fmt.Println(div.calculate(100, 2)) // = 50
}
