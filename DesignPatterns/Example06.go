package main

import (
	"math/rand"
	"time"
)

const size = 1000

var _ [size]int

func init() {
	rand.Seed(time.Now().UnixNano())

}

func main() {

}
