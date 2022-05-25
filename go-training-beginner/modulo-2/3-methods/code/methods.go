package main

import "fmt"

type Car struct {
	Name string
	Year int
}

func (c Car) IsLatest() bool {
	return c.Year >= 2018
}

// func (c Car) UpdateName(name string) {
// 	c.Name = name
// }

func (c *Car) UpdateName(name string) {
	c.Name = name
}

func main() {
	c := Car{"Tesla", 2021}

	fmt.Println("Is latest", c.IsLatest())

	c.UpdateName("Toyota")

	fmt.Println("Car:", c)
}
