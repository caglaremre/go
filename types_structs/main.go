package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Emre",
		LastName:  "Çağlar",
	}

	log.Println(user.FirstName, user.LastName, "phone number is ", user.PhoneNumber, "birthday is", user.BirthDate)
}
