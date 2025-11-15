package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return "Other direction"
	}
}

func main() {
	north := Direction(North)
	fmt.Println(north)

}
