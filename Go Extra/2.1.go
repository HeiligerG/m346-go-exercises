package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Gian",
			LastName:  "Where",
		},
		BirthDate: BirthDate{
			Day:   10,
			Month: 10,
			Year:  2000,
		},
		// TODO: set name and birth date information
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u264F', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)

	fmt.Printf("\nFull Profile:\n")
	fmt.Printf("Name: %s %s\n", me.Name.FirstName, me.Name.LastName)
	fmt.Printf("Birth Date: %02d-%02d-%d\n", me.BirthDate.Day, me.BirthDate.Month, me.BirthDate.Year)
	fmt.Printf("Siblings: %d\n", me.NumberOfSiblings)
	fmt.Printf("Zodiac Sign: %c\n", me.ZodiacSign)
}
