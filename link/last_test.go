package link

import (
	"fmt"
	"testing"
)

func TestDeleteLastN(t *testing.T) {
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

	ans := DeleteLastN(l1, 1)
	for ans != nil {
		fmt.Println(ans.Value)
		ans = ans.Next
	}
}

func TestDeleteLastN2(t *testing.T) {
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

	l1.Next = l2
	l2.Before = l1

	l2.Next = l3
	l3.Before = l2

	l3.Next = l4
	l4.Before = l3

	l4.Next = l5
	l5.Before = l4

	ans := DeleteLastN2(l1, 5)
	for ans != nil {
		before := "nil"
		if ans.Before != nil {
			before = fmt.Sprintf("%d", ans.Before.Value)
		}

		after := "nil"
		if ans.Next != nil {
			after = fmt.Sprintf("%d", ans.Next.Value)
		}

		fmt.Printf("%s<-%d->%s\n", before, ans.Value, after)

		ans = ans.Next
	}
}
