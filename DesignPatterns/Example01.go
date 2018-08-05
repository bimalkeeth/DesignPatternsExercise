package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Print(assert("rer", "rere"))
	fmt.Println("Hello,", Cyan("Aurora").Bold())
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
