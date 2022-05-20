package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type RostedNuts string
type Salad string

func (rn RostedNuts) PrepareDish() {
	fmt.Println("Rost your nuts")
}

func (s Salad) PrepareDish() {
	fmt.Println("Salad chopper")
	fmt.Println("Add papermint")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Preparing dishes:...")

	for _, dish := range dishes {
		fmt.Println("----- Dish : ", dish)
		dish.PrepareDish()
	}
}

const (
	Motorcycle = iota + 1
	Car
	Truck
)

const (
	smallLift = iota + 1
	standardLift
	largeLift
)

type Vehicle struct {
	model       string
	vehicleType uint
}

type MotorcycleVehicle Vehicle
type CarVehicle Vehicle
type TruckVehicle Vehicle

type Lifter interface {
	Lift()
}

func (m MotorcycleVehicle) Lift() {
	fmt.Println("Motorcycle", m, "Lift ", smallLift)
}

func (c CarVehicle) Lift() {
	fmt.Println("Car", c, "Lift ", standardLift)
}

func (t TruckVehicle) Lift() {
	fmt.Println("Truck", t, "Lift ", largeLift)
}

func main() {
	salad := Salad("Salad")
	rostedNuts := RostedNuts("RostedNuts")

	dishes := []Preparer{salad, rostedNuts}

	prepareDishes(dishes)

	//--Summary:
	//  Create a program that directs vehicles at a mechanic shop
	//  to the correct vehicle lift, based on vehicle size.
	//
	//--Requirements:
	//* The shop has lifts for multiple vehicle sizes/types:
	//  - Motorcycles: small lifts
	//  - Cars: standard lifts
	//  - Trucks: large lifts
	//* Write a single function to handle all of the vehicles
	//  that the shop works on.
	//* Vehicles have a model name in addition to the vehicle type:
	//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	//
	//--Notes:
	//* Use any names for vehicle models

	vehicles := []Lifter{
		MotorcycleVehicle{model: "MC 1", vehicleType: Motorcycle},
		MotorcycleVehicle{model: "MC 2", vehicleType: Motorcycle},
		CarVehicle{model: "Car 1", vehicleType: Car},
		CarVehicle{model: "Car 2", vehicleType: Car},
		TruckVehicle{model: "Truck 1", vehicleType: Truck},
		TruckVehicle{model: "Truck 2", vehicleType: Truck},
	}

	for _, vehicle := range vehicles {
		vehicle.Lift()
	}
}
