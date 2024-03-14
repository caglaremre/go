package main

import "fmt"

func main() {
	fmt.Println("hello, world")
	var whatToSay string
	var i int

	whatToSay = "goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingWasSaid := saySomething()

	fmt.Println("function returned", whatWasSaid, theOtherThingWasSaid)

}

func saySomething() (string, string) {
	return "say something", "else"
}
