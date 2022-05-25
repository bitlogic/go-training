package main

import "fmt"

type Pet struct {
	Name string
}

func main() {
	var m = map[string]Pet{
		"a": {"Gopher"},
		"b": {"Hedwing"},
	}

	//add values
	m["c"] = Pet{"Bito"}

	fmt.Println(m)

	//retrieve
	c := m["c"]

	fmt.Println("Key c:", c)

	//exist
	c, ok := m["c"]
	fmt.Println("Key c:", c, ok)

	d, ok := m["d"]
	fmt.Println("Key d:", d, ok)

	//update value
	m["a"] = Pet{"Chita"}
	fmt.Println(m)

	//delete value
	delete(m, "b")
	fmt.Println(m)

	//iteration

	for key, value := range m {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	//----//
	fmt.Println()
	//----//

	//properties

	m2 := m
	m2["c"] = Pet{"Bito"}

	fmt.Println(m)
	fmt.Println(m2)

}
