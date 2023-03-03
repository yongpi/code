package link

import (
	"fmt"
	"testing"
)

func TestTwoMeet(t *testing.T) {
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
		Value: 1,
		Next: &Link{
			Value: 3,
			Next: &Link{
				Value: 4,
				Next: &Link{
					Value: 5,
				},
			},
		},
	}

	fmt.Println(TwoMeet(l1, l2))
}

func TestCircleMeet(t *testing.T) {
	l1 := &Link{
		Value: 1,
	}

	l2 := &Link{
		Value: 2,
	}

	l3 := &Link{
		Value: 3,
	}

	l4 := &Link{
		Value: 4,
	}

	l5 := &Link{
		Value: 5,
	}

	l6 := &Link{
		Value: 6,
	}

	l7 := &Link{
		Value: 7,
	}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l7.Next = l4

	fmt.Println(CircleMeet(l1))
}
