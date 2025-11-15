package main

import (
	"fmt"
)

type Room struct {
	Name  string
	Clean bool
}

func checkRoomIsCleaned(rooms [5]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.Clean {
			fmt.Println(room.Name, "is clean")
		} else {
			fmt.Println(room.Name, "is not clean")
		}
	}
}

func main() {
	room1 := Room{"Agege", false}
	room2 := Room{"Shomolu", true}
	room3 := Room{"Epe", true}
	room4 := Room{"Isolo", false}
	room5 := Room{"Surulere", true}
	hotelRooms := [...]Room{room1, room2, room3, room4, room5}
	checkRoomIsCleaned(hotelRooms)
	for i := 0; i < len(hotelRooms); i++ {
		hotelRooms[i].Clean = false
	}
	checkRoomIsCleaned(hotelRooms)
	new()
	mapDemo()

}
func new() {
	hotelRooms := [...]Room{
		{"Agege", false},
		{"Shomolu", true},
		{"Epe", true},
		{"Isolo", false},
		{"Surulere", true},
	}
	checkRoomIsCleaned(hotelRooms)
}
