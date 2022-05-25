package main

import "fmt"

type Pet struct {
	Name string
}

func main() {
	var m map[string]int

	fmt.Println(m)

	if m == nil {
		fmt.Println("nil is zero value to maps")
	}

	//----//
	fmt.Println()
	//----//

	//Initialization

	//make
	var m1 = make(map[string]int)

	fmt.Println(m1)

	//literal

	var m3 = map[string]int{
		"a": 0,
		"b": 1,
	}

	fmt.Println(m3)

	var m4 = map[string]Pet{
		"a": Pet{"Gopher"},
		"b": Pet{"Hedwing"},
	}

	/*
		var m4 = map[string]Pet{
			"a": {"Gopher"},
			"b": {"Hedwing"},
		}
	*/

	fmt.Println(m4)

}
