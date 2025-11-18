//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
//

package main

import (
	"testing"
)

func newPlayer() Player {
	return Player{
		name:      "Warrior",
		Health:    100,
		MaxHealth: 100,
		Energy:    100,
		MaxEnergy: 100,
	}
}

func TestApplyDamage(t *testing.T) {
	player := newPlayer()
	player.addHealth(999)
	if player.Health > player.MaxHealth {
		t.Fatalf("Health went over limit :%v, want %v", player.Health, player.MaxHealth)
	}

	player.applyDamage(player.MaxHealth + 1)
	if player.Health < 0 {
		t.Fatalf("Health: %v. Minimum: o", player.Health)
	}
	if player.Health > player.MaxHealth {
		t.Fatalf("health : %v. maximum: %v", player.Health, player.MaxHealth)
	}
}

func TestConsumeEnergy(t *testing.T) {
	player := newPlayer()
	player.consumeEnergy(999)
	if player.Energy > player.MaxEnergy {
		t.Fatalf("Energy went over limit :%v, want %v", player.Energy, player.MaxEnergy)
	}

	player.consumeEnergy(player.Energy + 1)
	if player.Energy < 0 {
		t.Fatalf("Energyh: %v. Minimum: o", player.Energy)
	}
	if player.Energy > player.MaxEnergy {
		t.Fatalf("Energy : %v. maximum: %v", player.Energy, player.MaxEnergy)
	}
}
