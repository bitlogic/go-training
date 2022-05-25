package main

import "fmt"

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	i := MyInt(10)

	fmt.Println(i.isGreater(5))
}
