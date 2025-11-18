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

type Energy uint
type Health uint
type Name string

type Player struct {
	name              Name
	Health, MaxHealth Health
	Energy, MaxEnergy Energy
}

func (player *Player) applyDamage(loss Health) {
	if player.Health-loss > player.Health {
		player.Health = 0
	} else {
		player.Health -= loss
	}
	fmt.Println(player.name, "Damage", loss, "health->", player.Health)

}
func (player *Player) addHealth(add Health) {
	player.Health += add
	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}
	fmt.Println(player.name, "Add", add, "health ->", player.Health)
}
func (player *Player) consumeEnergy(loss Energy) {
	if player.Energy-loss > player.Energy {
		player.Energy = 0
	} else {
		player.Energy -= loss
	}
	fmt.Println(player.name, "Consume", loss, "energy->", player.Energy)
}

func (player *Player) addEnergy(add Energy) {
	player.Energy += add
	if player.Energy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	}
	fmt.Println(player.name, "Add", add, "health ->", player.Energy)
}

func main() {
	player := Player{
		name:      "Knight",
		Health:    100,
		MaxHealth: 100,
		Energy:    500,
		MaxEnergy: 500,
	}
	player.applyDamage(44)
	player.addHealth(12)
	player.consumeEnergy(9999)
	player.addEnergy(20)
	fmt.Println(player)
}
