package main

import "fmt"

func main() {
	// var a = [4]int{1, 2, 3, 4}
	// var b [2]int = a // error: cannot use a (variable of type [4]int) as type [2]int in variable declaration

	var a = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var b = a // Copy of a is assigned to b

	b[0] = "Monday"

	fmt.Println(a)
	fmt.Println(b)
}
