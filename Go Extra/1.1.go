package main

import "fmt"

func main() {

	var firstName, lastName = "James", "Bond"
	var dayOfBirth, monthOfBirth, yearOfBirth, numberOfSiblings = 10, 10, 2000, 2
	var heightInMeters = 1.8
	var zodiacSign = '\u264F'

	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
