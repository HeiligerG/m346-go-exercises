package main

import "fmt"

var lastName string = "Last"

var firstName string = "First"

var dayOfBirth int = 12

var monthOfBirth int = 10

var yearOfBirth int = 2000

var numberOfSiblings int = 3

var heightInMeters float32 = 1.2

var zodiacSign string = "Scorpion"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
