package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			firstName: "Erion",
			lastName:  "Jasiqi",
		},
		BirthDate: BirthDate{
			dayOfBirth:   4,
			monthOfBirth: 9,
			yearOfBirth:  2007,
		},
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u264D', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister

	fmt.Println("Siblings After:", me.NumberOfSiblings+1)
}
