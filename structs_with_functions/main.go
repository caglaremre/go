package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) PrintFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Emre"

	myVar2 := myStruct{
		FirstName: "Ege",
	}

	log.Println("myVar is set to", myVar.PrintFirstName())
	log.Println("myVar2 is set to", myVar2.PrintFirstName())
}
