package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("First")

	panicker(0)

	fmt.Println("Last")
}

func panicker(num int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Tried to recover from error", r)
		}
	}()


	if sum := 0 / num; sum != 0 {
		fmt.Println("You can't really divide by Zer0")
	}
}
