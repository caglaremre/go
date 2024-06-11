package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cat", "cow"}
	log.Println("#################")

	// "_" blank identifier for not used variable for compiler
	for _, animal := range animals {
		log.Println(animal)
	}
	log.Println("#################")

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"

	for animalType, animal := range animalsMap {
		log.Println(animalType, animal)
	}
	log.Println("#################")

	var firstLine = "one ring to rule them all"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
}
