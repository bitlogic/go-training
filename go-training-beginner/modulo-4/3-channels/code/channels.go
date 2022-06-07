package main

import "fmt"

func speak(arg string, ch chan string) {
	ch <- arg // Send
}

func main() {
	ch := make(chan string)

	go speak("Hello World", ch)

	data := <-ch // Receive
	fmt.Println(data)
}
