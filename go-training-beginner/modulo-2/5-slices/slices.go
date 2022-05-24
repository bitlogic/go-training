package main

import "fmt"

func main() {
	array := [5]int{20, 15, 5, 30, 25}

	slice := array[1:4]

	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", array, len(array), cap(array))

	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", slice, len(slice), cap(slice))

	//----//
	fmt.Printf("\n")
	//----//

	var s []string

	fmt.Println(s)
	fmt.Println(s == nil)

	//----//
	fmt.Printf("\n")
	//----//

	//initialize
	var a = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s1 := a[0:2] // Select from 0 to 2
	s2 := a[:3]  // Select first 3
	s3 := a[2:]  // Select last 2

	fmt.Println("Array:", a)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)

	//----//
	fmt.Printf("\n")
	//----//

	//functions

	//copy

	s4 := []string{"a", "b", "c", "d"}
	s5 := make([]string, 4)

	e := copy(s5, s4)

	fmt.Println("Src:", s4)
	fmt.Println("Dst:", s5)
	fmt.Println("Elements:", e)

	//append

	s6 := []string{"a", "b", "c", "d"}

	s7 := append(s6, "e", "f")

	fmt.Println("\ns6:", s6)
	fmt.Println("s7:", s7)

	//----//
	fmt.Printf("\n")
	//----//

	//Properties

	array2 := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	slice2 := array2[0:2]

	fmt.Println(array2)
	fmt.Println(slice2)

	slice2[0] = "Sun"

	fmt.Println(slice2)

	//----//
	fmt.Printf("\n")
	//----//

	// variadic function
	values := []int{1, 2, 3}
	sum := add(values...)
	fmt.Println(sum)

}

func add(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}
