package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	histgrams := make(map[string]int)
	words := words(data)
	for word := range words {
		histgrams[word]++
	}
	for k, v := range histgrams {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
func words(data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				out <- word
			}
		}
	}()
	return out
}
