package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	Classes := map[string]Class{
		"D23a": {
			Students: []Student{
				{FirstName: "Hans", LastName: "Ruedi"},
				{FirstName: "Bob", LastName: "Bauer"},
			},
		},
		"D23b": {
			Students: []Student{
				{FirstName: "Max", LastName: "Muster"},
				{FirstName: "Jonas", LastName: "Meister"},
			},
		},
	}

	// TODO: output everything using fmt.Println()
	for className, class := range Classes {
		fmt.Printf(" %s:\n", className)
		fmt.Println("Students:")
		for _, student := range class.Students {
			fmt.Printf("%s %s\n", student.FirstName, student.LastName)
		}
		fmt.Println()
	}

}
