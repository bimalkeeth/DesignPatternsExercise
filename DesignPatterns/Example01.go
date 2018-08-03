package main

import "fmt"

func main() {
	fmt.Print(assert("rer", "rere"))
}
func assert(c string, s string) bool {
	switch name, surname := "rer", "rere"; {
	case c == name:
		return true
	case s == surname:
		return true
	}
	return false
}
