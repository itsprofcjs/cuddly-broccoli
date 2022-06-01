package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) freeSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

type Health struct {
	value     int
	maxHealth int
}

type Energy struct {
	value     int
	maxEnergy int
}

type Player struct {
	health Health
	energy Energy
	name   string
}

func (player *Player) changeHealth(newLevel int) {
	player.health.value = newLevel

	fmt.Println(player)
}

func (player *Player) changeEnergy(newLevel int) {
	player.energy.value = newLevel

	fmt.Println(player)
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println(lot)

	lot.occupySpace(1)
	occupySpace(&lot, 4)
	fmt.Println(lot)

	lot.freeSpace(4)
	fmt.Println(lot)

	player1 := Player{name: "CJS"}

	player1.changeEnergy(5)
	player1.changeHealth(80)
}
