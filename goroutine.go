package main

import (
	"fmt"
	"math"
)

func main() {
	resultChan := make(chan float64)

	go func() {
		resultChan <- math.Sqrt(4)
	}()

	result := <-resultChan
	fmt.Println(result)
}
