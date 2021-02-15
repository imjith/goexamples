package main

import (
	"fmt"
	"time"
)

func display() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Inside display()")
	}
}

func main() {
	go display()

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Inside main()")
	}
}
