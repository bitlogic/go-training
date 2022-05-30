package main

import "fmt"

func main() {
	WillPanic()

	fmt.Println("Calm down, everything will be fine")
}

func handlePanic() {
	data := recover()
	fmt.Println("Recovered:", data)
}

func WillPanic() {
	defer handlePanic()

	panic("Ahhh!!!!")
}
