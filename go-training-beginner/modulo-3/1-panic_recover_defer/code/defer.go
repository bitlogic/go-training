package main

import "fmt"

func main() {
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you sure?") //last input

	fmt.Println("Doing some work...")
}
