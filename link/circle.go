package link

func CircleDelete(head *Link, m int) *Link {
	if head == nil || head.Next == nil {
		return head
	}

	last := head
	for last.Next != head {
		last = last.Next
	}

	var count int
	for last != head {
		count++
		if count == m {
			count = 0
			last.Next = head.Next
		} else {
			last = last.Next
		}

		head = head.Next
	}

	return head
}
