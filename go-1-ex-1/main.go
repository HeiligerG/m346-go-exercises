package main

import "fmt"

var lastName, firstName = "Last", "First"

var monthOfBirth, yearOfBirth, dayOfBirth, numberOfSiblings = 10, 2000, 12, 3

var heightInMeters float32 = 1.2

var zodiacSign = '\u264F'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
