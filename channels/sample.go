package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	fmt.Println("Sending Data")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Sent %d\n", i)
		ch <- i
	}
	close(ch)
}

func receiveData(ch chan int) {
	fmt.Println("Receiving data")
	for {
		n, ok := <-ch
		if ok {
			fmt.Printf("Received %d\n", n)
		} else {
			fmt.Println("Channel closed")
			break
		}
	}

}

func main() {
	ch := make(chan int)

	go sendData(ch)
	go receiveData(ch)

	time.Sleep(5 * time.Second) // wait for the functions to complete
	fmt.Println("Exiting main...")
}
