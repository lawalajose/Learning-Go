package main

import "fmt"

type Space struct {
	occupied bool
}
type ParkingLot struct {
	space []Space
}

func occupySpace(lot *ParkingLot, number int) {
	lot.space[number-1].occupied = true
}

func (lot *ParkingLot) occupySpaces(number int) {
	lot.space[number-1].occupied = true
}
func (lot ParkingLot) occupySpacess(number int) {
	lot.space[number-1].occupied = true
}

func main() {
	lot := ParkingLot{make([]Space, 10)}

	occupySpace(&lot, 1)
	lot.occupySpaces(2)
	lot.occupySpacess(3)
	for k, v := range lot.space {
		fmt.Println("Space:", k, "ocuupied _", v)
	}

}
