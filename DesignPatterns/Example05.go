package main

import "fmt"

func main() {
	seven := [7]string{"grumpy", "sleepy", "bashful"}
	fmt.Println(len(seven), cap(seven))
}
