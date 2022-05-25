package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)

	fmt.Println(s)

	s1, ok := i.(string)

	fmt.Println(s1, ok)

	f, ok := i.(float64)

	fmt.Println(f, ok)

	// f = i.(float64)

	// fmt.Println(f)

	//----//
	fmt.Println()
	//----//

	//switch

	var t interface{}

	fmt.Println(t)

	t = 100

	switch t := t.(type) {
	case string:
		fmt.Printf("string: %s\n", t)
	case bool:
		fmt.Printf("boolean: %v\n", t)
	case int:
		fmt.Printf("integer: %d\n", t)
	default:
		fmt.Printf("unexpected: %T\n", t)
	}
}
