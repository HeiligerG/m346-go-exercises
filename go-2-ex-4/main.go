package main

import "fmt"

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}

	type Class struct {
		Students []Student
	}

	modules := map[string]Class{
		"Math101": {
			Students: []Student{
				{FirstName: "Alice", LastName: "Smith"},
				{FirstName: "Bob", LastName: "Johnson"},
			},
		},
		"Physics201": {
			Students: []Student{
				{FirstName: "Charlie", LastName: "Brown"},
				{FirstName: "Diana", LastName: "White"},
			},
		},
	}

	for moduleName, class := range modules {
		fmt.Println("Module:", moduleName)
		for _, student := range class.Students {
			fmt.Printf("  - %s %s\n", student.FirstName, student.LastName)
		}
	}
}

// TODO: declare a type for Student (with first and last name)
// TODO: declare a type for Class (consisting of multiple students)
// TODO: declare a map of modules being attended by multiple classes
// TODO: output everything using fmt.Println()
