//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

type rectangle struct {
	lenght, width int
}

func calculateRectangle(lenght, width int) {
	area := lenght * width
	fmt.Println("The area of this rectangle of lenght", lenght, "and width", width, "is", area)
}
func calculatePerimeter(lenght, width int) {
	perimeter := (lenght * 2) + (width * 2)
	fmt.Println("The perimeter of this rectangle of lenght", lenght, "and width", width, "is", perimeter)
}

func main() {
	rectangle1 := rectangle{
		lenght: 5,
		width:  8,
	}
	calculateRectangle(rectangle1.lenght, rectangle1.width)
	calculatePerimeter(rectangle1.lenght, rectangle1.width)
	rectangle1.lenght = rectangle1.lenght * 2
	rectangle1.width = rectangle1.width * 2
	calculateRectangle(rectangle1.lenght, rectangle1.width)
	calculatePerimeter(rectangle1.lenght, rectangle1.width)

}
