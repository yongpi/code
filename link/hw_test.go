package link

import (
	"fmt"
	"testing"
)

func TestIsHw(t *testing.T) {
	l1 := &Link{
		Value: 1,
		Next: &Link{
			Value: 2,
			Next: &Link{
				Value: 2,
				Next: &Link{
					Value: 1,
				},
			},
		},
	}

	fmt.Println(IsHw(l1))
}

func TestIsHw2(t *testing.T) {
	l1 := &Link{
		Value: 1,
		Next: &Link{
			Value: 2,
			Next: &Link{
				Value: 3,
				Next: &Link{
					Value: 2,
					Next: &Link{
						Value: 1,
					},
				},
			},
		},
	}

	ok, head := IsHw2(l1)
	fmt.Println(ok)
	fmt.Println(head)
}
