package main

import "fmt"

type matrix [2][2][2][2]byte

func main() {
	var mat1 matrix
	mat1 = initMat()
	fmt.Println(mat1)
}
func initMat() matrix {

	return matrix{
		{{{4, 2}, {5, 7}}, {{9, 6}, {5, 4}}},
		{{{5, 9}, {3, 8}}, {{8, 6}, {3, 4}}},
	}

}
