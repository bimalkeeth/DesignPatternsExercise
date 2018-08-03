package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}
