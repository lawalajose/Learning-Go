package main

import (
	"fmt"
)

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Frontseat struct {
	Passenger Passenger
}

func main() {
	Cassey := Passenger{
		Name:         "Cassey",
		TicketNumber: 1,
	}
	fmt.Println(Cassey)
	Bill := Passenger{"Bill", 2, false}
	Ella := Passenger{"Ella", 3, false}
	fmt.Println(Bill, Ella)
	Heidi := Passenger{
		Name:         "Heidi",
		TicketNumber: 4,
	}
	fmt.Println(Heidi)
	Bill.Boarded = true
	Cassey.Boarded = true
	if Bill.Boarded {
		fmt.Println(Bill.Name, "has boarded the bus")
	}
	if Cassey.Boarded {
		fmt.Println(Cassey.Name, "has boarded the bus")
	}
	Heidi.Boarded = true
	NewHeidi := Frontseat{
		Passenger: Heidi,
	}
	fmt.Println(NewHeidi)
	fmt.Println(NewHeidi.Passenger.Name, "is in front seat")

}
