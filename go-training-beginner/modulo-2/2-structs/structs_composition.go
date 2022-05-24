package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

type SuperHero struct {
	RealName Person
	Alias    string
}

func main() {
	p := Person{"Bruce", "Wayne", 40}
	s := SuperHero{p, "batman"}

	fmt.Println(s)
}
