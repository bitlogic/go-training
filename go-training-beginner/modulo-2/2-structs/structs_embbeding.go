package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

type SuperHero struct {
	Person
	Alias string
}

func main() {
	s := SuperHero{}

	s.FirstName = "Bruce"
	s.LastName = "Wayne"
	s.Age = 40
	s.Alias = "batman"

	fmt.Println(s)
}
