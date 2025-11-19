package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index %v", index))
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{
		values: []int{1, 3, 4, 5},
	}

	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value is %v ", value)
	}
	fmt.Println()
}
