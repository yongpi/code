package link

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	l1 := &Link{
		Value: 1,
		Next: &Link{
			Value: 2,
			Next: &Link{
				Value: 3,
				Next: &Link{
					Value: 4,
					Next: &Link{
						Value:  5,
						Next:   nil,
						Before: nil,
					},
				},
			},
		},
	}

	l2 := &Link{
		Value: 2,
		Next: &Link{
			Value: 4,
			Next: &Link{
				Value: 8,
				Next: &Link{
					Value: 10,
					Next: &Link{
						Value: 12,
					},
				},
			},
		},
	}

	ans := MergeSort(l1, l2)
	for ans != nil {
		fmt.Println(ans.Value)
		ans = ans.Next
	}
}
