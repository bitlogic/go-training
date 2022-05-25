package main

import "fmt"

func main() {
	var p *int

	fmt.Println(p)

	//-----//

	a := 10

	var p1 *int = &a

	fmt.Println("address p1:", p1)

	//-----//

	p2 := new(int)
	*p2 = 100

	fmt.Println("value:", *p2)
	fmt.Println("address p2:", p2)

	//-----//

	fmt.Println("before", a)
	fmt.Println("address p1:", p1)

	*p1 = 20
	fmt.Println("after:", a)
	fmt.Println("address:", p1)
}
