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
	Health, MaxHealth uint // u int is unsigned integer
	Energy, MaxEnergy uint
	Name              string
}

// receiver functions to modify player stats
func (player *Player) modifyHealth(newHealth uint) {
	player.Health = newHealth
}

func (player *Player) modifyEnergy(newEnergy uint) {
	player.Energy = newEnergy
}

func main() {
	john := Player{20, 40, 30, 50, "John"}
	fmt.Println("Initial:")
	fmt.Println("Health:", john.Health, "Max Health:", john.MaxHealth, "Energy:", john.Energy, "Max Energy:", john.MaxEnergy)
	john.modifyHealth(30)
	fmt.Println("After updating health:")
	fmt.Println("Health:", john.Health, "Max Health:", john.MaxHealth, "Energy:", john.Energy, "Max Energy:", john.MaxEnergy)
	john.modifyEnergy(35)
	fmt.Println("After updating energy:")
	fmt.Println("Health:", john.Health, "Max Health:", john.MaxHealth, "Energy:", john.Energy, "Max Energy:", john.MaxEnergy)
}
