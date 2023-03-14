package tree

import (
	"fmt"
	"testing"
)

func TestMaxValue(t *testing.T) {
	root := &Tree{
		Left: &Tree{
			Left: &Tree{
				Value: 4,
			},
			Right: &Tree{
				Value: 5,
			},
			Value: 2,
		},
		Right: &Tree{
			Left: &Tree{
				Value: 6,
			},
			Right: &Tree{
				Value: 6,
			},
			Value: 3,
		},
		Value: 1,
	}

	fmt.Println(MaxValue(root))
}
