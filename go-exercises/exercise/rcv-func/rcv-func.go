//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Energy int
type Health int
type Name string

type Player struct {
	name      Name
	maxHealth Health
	maxEnergy Energy
}

func (player *Player) energyDeduction(energyloss Energy) {
	player.maxEnergy -= energyloss
}

func (player *Player) healthDeduction(healthloss Health) {
	player.maxHealth -= healthloss
}

func main() {
	player1 := Player{"player-1", 100, 50}
	fmt.Println(player1)
	player1.energyDeduction(30)
	player1.healthDeduction(10)
	fmt.Println(player1)

}
