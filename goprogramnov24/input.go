package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	fmt.Print("Enter name for person 1: ")
	fmt.Scanln(&pers1.name)
	fmt.Print("Enter age for person 1: ")
	fmt.Scanln(&pers1.age)
	fmt.Print("Enter job for person 1: ")
	fmt.Scanln(&pers1.job)
	fmt.Print("Enter salary for person 1: ")
	fmt.Scanln(&pers1.salary)

	// Access and print Pers1 info
	fmt.Println("\nPerson 1 Info:")
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)
}
