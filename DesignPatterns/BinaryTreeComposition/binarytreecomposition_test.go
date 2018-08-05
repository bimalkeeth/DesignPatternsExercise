package binarytreecomposition

import (
	"fmt"
	"testing"
)

func TestComposition(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{LeafValue: 6},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}

	fmt.Println(root.Right.LeafValue)
}
