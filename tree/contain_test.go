package tree

import (
	"fmt"
	"testing"
)

func TestContain(t *testing.T) {
	node1 := &Tree{
		Left: &Tree{
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
		},
		Right: &Tree{
			Left: &Tree{
				Value: 6,
			},
			Right: &Tree{
				Value: 7,
			},
			Value: 3,
		},
		Value: 1,
	}

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

	fmt.Println(Contain(node1, node2))
}
