package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("before", a)
	fmt.Println("address p1:", p)

	myFunction(&a) // or myFunction(p)

	fmt.Println("after:", a)
	fmt.Println("address:", p)

	p1 := &a

	fmt.Println(p == p1)

}

func myFunction(ptr *int) {
	*ptr = 20
}
