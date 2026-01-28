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

type Player struct {
	Name      string
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
}

func (player *Player) setMaxHealth(amount int) {
	oldHealth := player.Health
	player.Health += amount

	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}

	if player.Health < 0 {
		player.Health = 0
	}

	fmt.Printf(
		"%s ning Health o‘zgardi: %d -> %d (farq: %+d)\n",
		player.Name, oldHealth, player.Health, amount,
	)

}

func (player *Player) setMaxEnergy(amount int) {
	oldEnergy := player.Energy
	player.Energy += amount
	if player.Energy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	}
	if player.Energy < 0 {
		player.Energy = 0
	}

	fmt.Printf(
		"%s ning Health o‘zgardi: %d -> %d (farq: %+d)\n",
		player.Name, oldEnergy, player.Energy, amount,
	)
}

func main() {

	player := Player{
		Name:      "Russel",
		Health:    34,
		MaxHealth: 90,
		Energy:    31,
		MaxEnergy: 85,
	}

	player.setMaxHealth(10)
	player.setMaxEnergy(50)
}
