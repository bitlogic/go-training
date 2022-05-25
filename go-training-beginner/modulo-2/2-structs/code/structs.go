package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p1 Person

	fmt.Println("Person 1:", p1)

	var p2 = Person{
		FirstName: "Karan",
		LastName:  "Pratap Singh",
		Age:       22,
	}

	fmt.Println("Person 2:", p2)

	var p3 = Person{
		FirstName: "Tony",
		LastName:  "Stark",
	}

	fmt.Println("Person 3:", p3)

	var p4 = Person{"Bruce", "Wayne", 40}

	fmt.Println("Person 4:", p4)

	var a = struct {
		Name    string
		IsGreat bool
	}{"Golang", true}

	b := struct {
		Name    string
		IsGreat bool
	}{"PHP", false}

	fmt.Println("Anonymous:", a)
	fmt.Println("Anonymous:", b)

	//-------//

	p5 := new(Person)

	p5.FirstName = "Clark"
	p5.LastName = "Ken"
	p5.Age = 22

	fmt.Println("Person", p5)

	//-------//
}
