package main

import "log"

func main() {
	var myString string
	myString = "green"

	log.Println("myString is", myString)
	changeUsingPointer(&myString)
	log.Println("after using changeUsingPointer my String is", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "red"
	*s = newValue
}
