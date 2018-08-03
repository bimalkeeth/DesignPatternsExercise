package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	makeEvents(4, ch)
	for i := 0; i < 4; i++ {
		if val, opened := <-ch; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed")
		}
	}

}

func makeEvents(count int, in chan<- int) {
	for i := 0; i < count; i++ {
		in <- 2 * i
	}
	defer close(in)

}
