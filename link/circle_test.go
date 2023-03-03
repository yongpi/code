package link

import (
	"fmt"
	"testing"
)

func TestCircleDelete(t *testing.T) {
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
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l1

	fmt.Println(CircleDelete(l1, 1).Value)

}

func TestCircleInsert(t *testing.T) {
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
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l1

	ans := CircleInsert(l1, 4)
	for i := 0; i < 10; i++ {
		fmt.Println(ans.Value)
		ans = ans.Next
	}
}
