package link

import "fmt"

type Link struct {
	Value                int
	Next, Before, Random *Link
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

func DeleteSameValue(head *Link) *Link {
	hm := make(map[int][]*Link)

	node := head
	for node != nil {
		hm[node.Value] = append(hm[node.Value], node)
		node = node.Next
	}

	node = head
	for node.Next != nil {
		next := node.Next
		if data, ok := hm[next.Value]; ok && len(data) > 0 && data[0] == next {
			node = next
			continue
		}

		node.Next = node.Next.Next
	}

	return head
}
