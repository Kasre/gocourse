package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Printf("Goroutine: %v \n", n)
		}
	}()

	for n := range c {
		fmt.Printf("Main: %v \n", n)
	}
}
