package tree

import (
	"fmt"
	"testing"
)

func TestBalance(t *testing.T) {
	node2 := &Tree{
		Left: &Tree{
			Left: nil,
			Right: &Tree{
				Value: 8,
			},
			Value: 4,
		},
		Right: &Tree{
			Left: &Tree{
				Value: 9,
			},
			Value: 5,
		},
		Value: 2,
	}

	fmt.Println(Balance(node2))
}
