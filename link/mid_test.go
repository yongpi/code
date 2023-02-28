package link

import (
	"fmt"
	"testing"
)

func TestDeleteMid(t *testing.T) {
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

	ans := DeleteMid(l1)
	for ans != nil {
		fmt.Println(ans.Value)
		ans = ans.Next
	}
}
