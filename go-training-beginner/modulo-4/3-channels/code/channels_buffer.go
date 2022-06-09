package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Hello World"
	ch <- "Hi again"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
