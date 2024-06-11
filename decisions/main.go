package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	Cat := "cat2"

	if Cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isFalse := false

	if myNum > 99 && !isFalse {
		log.Println("myNum is greater than 99 and isFalse is set to true")
	}

	// switch statement

	myVar := "fish"
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "fish":
		log.Println("myVar is set to fish")
	default:
		log.Println("myVar is someting else")
	}

}
