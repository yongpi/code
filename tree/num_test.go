package tree

import (
	"fmt"
	"testing"
)

func TestFindExchangeNode(t *testing.T) {
	root := &Tree{
		Left: &Tree{
			Left: &Tree{
				Value: 1,
			},
			Value: 2,
		},
		Right: &Tree{
			Right: &Tree{
				Value: 4,
			},
			Value: 5,
		},
		Value: 3,
	}

	for _, value := range FindExchangeNode(root) {
		fmt.Println(value.Value)
	}
}
