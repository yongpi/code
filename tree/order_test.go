package tree

import (
	"testing"
)

func TestPostOrder(t *testing.T) {
	var orderTree = &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 3,
			},
			Right: &Tree{
				Value: 4,
			},
		},
		Right: &Tree{
			Value: 5,
			Left: &Tree{
				Value: 6,
			},
			Right: &Tree{
				Value: 7,
			},
		},
	}

	PostOrder(orderTree)
}
