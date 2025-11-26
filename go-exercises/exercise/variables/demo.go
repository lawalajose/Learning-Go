package main

import (
	"fmt"
)

type Phone struct {
	Name  string
	Brand string
	Cost  int
}

func (p Phone) String() string {
	return fmt.Sprintf("The %v made by %v costs %d naira", p.Name, p.Brand, p.Cost)
}

func main() {
	ISPM := Phone{
		Name:  "Iphone 17 pro max",
		Brand: "Apple",
		Cost:  2540000,
	}
	fmt.Println(ISPM)

}
