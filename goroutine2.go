package main

import (
	"fmt"
	"time"
)

func main() {
	Aspak := make(chan string)
	Prasant := make(chan string)

	go func() {
		for {
			msg := <-Prasant
			fmt.Println("Aspak received:", msg)
			fmt.Println("Aspak sending: Hello, Prasant!")
			Aspak <- "Hello, Prasant!"
		}
	}()

	go func() {
		for {
			msg := <-Aspak
			fmt.Println("Prasant received:", msg)
			fmt.Println("Prasant sending: Hello, Aspak!")
			Prasant <- "Hello, Aspak!"
		}
	}()
	fmt.Println("Start Messaging")
	Aspak <- "Hello, Prasant!"
	time.Sleep(time.Second)
}
