package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name           string
	Color          string
	NumberOfTheeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephard",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:           "Jack",
		Color:          "Silver",
		NumberOfTheeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal say", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "unga bunga"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
