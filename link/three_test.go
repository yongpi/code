package link

import (
	"fmt"
	"testing"
)

func TestMakeThree(t *testing.T) {
	l1 := &Link{
		Value: 1,
		Next: &Link{
			Value: 5,
			Next: &Link{
				Value: 2,
				Next: &Link{
					Value: 4,
					Next: &Link{
						Value: 3,
					},
				},
			},
		},
	}

	ans := MakeThree(l1, 3)
	for ans != nil {
		fmt.Println(ans.Value)
		ans = ans.Next
	}

}
