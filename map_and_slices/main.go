package main

import (
	"log"
	"slices"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "emre",
		LastName:  "çağlar",
	}

	myMap["me"] = me
	log.Println(myMap["me"].FirstName)

	var mySlice []string

	mySlice = append(mySlice, "emre")
	mySlice = append(mySlice, "ege")

	log.Println(mySlice)

	var mySecondSlice []int

	mySecondSlice = append(mySecondSlice, 2)
	mySecondSlice = append(mySecondSlice, 1)
	mySecondSlice = append(mySecondSlice, 3)

	//sort.Ints(mySecondSlice) --> old way
	slices.Sort(mySecondSlice)

	log.Println(mySecondSlice)

	// different way to define slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)

}
