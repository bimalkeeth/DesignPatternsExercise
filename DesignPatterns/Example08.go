package main

import "fmt"

type food interface {
	eat()
}

type veggie string

func (v veggie) eat() {
	fmt.Println("Eating", v)
}

type meat string

func (m meat) eat() {
	fmt.Println("Eating Tasty", m)
}

func eat(f food) {
	veg, ok := f.(veggie)
	if ok {
		if veg == "okra" {
			fmt.Println("Yuk! not eating ", veg)
		} else {
			veg.eat()
		}
		return
	}
	mt, ok := f.(meat)
	if ok {
		if mt == "beef" {
			fmt.Println("Yuk! not eating ", mt)
		} else {
			mt.eat()
		}
		return
	}
	fmt.Println("Not eating whatever that is: ", f)
}

func eactnext(f food) {

	switch morsel := f.(type) {
	case veggie:
		if morsel == "okra" {
			fmt.Println("Yuk! not eating ", morsel)
		} else {
			morsel.eat()
		}
	case meat:
		if morsel == "beef" {
			fmt.Println("Yuk! not eating ", morsel)
		} else {
			morsel.eat()
		}
	default:
		fmt.Println("Not eating whatever that is: ", f)
	}

}

func main() {

	var s meat = "beef"
	eat(s)
	eactnext(s)
}
