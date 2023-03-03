package link

func MergeSort(l1, l2 *Link) *Link {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *Link
	for l1 != nil && l2 != nil {
		item := l1
		if l2.Value < item.Value {
			item = l2
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}

		item.Next = nil
		if head == nil {
			head = item
			node = item
		} else {
			node.Next = item
			node = item
		}
	}

	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return head
}
