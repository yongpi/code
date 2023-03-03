package link

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	l1 := &Link{
		Value: 0,
		Next: &Link{
			Value: 4,
			Next: &Link{
				Value: 1,
				Next: &Link{
					Value: 3,
					Next: &Link{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}

	ans := SelectionSort(l1)
	for ans != nil {
		fmt.Println(ans.Value)
		ans = ans.Next
	}
}
