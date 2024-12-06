package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	Name             FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "John",
			LastName:  "Doe",
		},
		BirthDate: BirthDate{
			Day:   15,
			Month: 7,
			Year:  1995,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       '\u264B',
	}
	fmt.Printf("Profile: %s %s, born %d.%d.%d, %d siblings, zodiac: %c\n",
		me.Name.FirstName, me.Name.LastName,
		me.BirthDate.Day, me.BirthDate.Month, me.BirthDate.Year,
		me.NumberOfSiblings, me.ZodiacSign)
}
