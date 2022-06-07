package main

import (
	"fmt"
	"time"
)

func speak(arg string) {
	fmt.Println(arg)
}

func main() {
	go speak("Hello World")

	time.Sleep(1 * time.Second)
}
