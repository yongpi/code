package link

import (
	"fmt"
	"testing"
)

func TestCopyRandom(t *testing.T) {
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

	l1.Random = l5
	l3.Random = l4

	ans := CopyRandom(l1)
	for ans != nil {
		before := "nil"
		if ans.Random != nil {
			before = fmt.Sprintf("%d", ans.Random.Value)
		}

		after := "nil"
		if ans.Next != nil {
			after = fmt.Sprintf("%d", ans.Next.Value)
		}

		fmt.Printf("%s<-%d->%s\n", before, ans.Value, after)

		ans = ans.Next
	}
}

func TestCopyRandom2(t *testing.T) {
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

	l1.Random = l5
	l3.Random = l4

	ans := CopyRandom2(l1)
	for ans != nil {
		before := "nil"
		if ans.Random != nil {
			before = fmt.Sprintf("%d", ans.Random.Value)
		}

		after := "nil"
		if ans.Next != nil {
			after = fmt.Sprintf("%d", ans.Next.Value)
		}

		fmt.Printf("%s<-%d->%s\n", before, ans.Value, after)

		ans = ans.Next
	}
}
