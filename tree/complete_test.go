package tree

import (
	"fmt"
	"testing"
)

func TestIsCBT(t *testing.T) {
	node2 := &Tree{
		Left: &Tree{
			Left: &Tree{
				Value: 8,
			},
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

	fmt.Println(IsCBT(node2))
}
