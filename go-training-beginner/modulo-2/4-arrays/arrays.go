package main

import "fmt"

func main() {
	var arr1 [4]int

	fmt.Println("Array 1: ", arr1)

	var arr2 = [4]int{1, 2, 3, 4}

	fmt.Println("Array 2: ", arr2)

	//index
	fmt.Println("\nfirst element:", arr2[0])

	//iteration

	fmt.Printf("\nmethod with len\n")

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Index: %d, Element: %d\n", i, arr2[i])
	}

	fmt.Printf("\nmethod with range\n")

	// method 2
	for i, e := range arr2 {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

	//matriz

	fmt.Printf("\n")
	matriz1 := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range matriz1 {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

	fmt.Printf("\n")
	matriz2 := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range matriz2 {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

}
