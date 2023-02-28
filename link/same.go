package link

import "fmt"

type Link struct {
	Value        int
	Next, Before *Link
}

func SamePart(l1, l2 *Link) {
	for l1 != nil && l2 != nil {
		if l1.Value > l2.Value {
			l2 = l2.Next
		}

		if l1.Value < l2.Value {
			l1 = l1.Next
		}

		if l1.Value == l2.Value {
			fmt.Println(l1.Value)
			l1 = l1.Next
			l2 = l2.Next
		}
	}
}
