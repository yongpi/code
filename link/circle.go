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

func CircleInsert(head *Link, value int) *Link {
	end := head
	for end.Next != head {
		end = end.Next
	}

	if value > head.Value && value < end.Value {
		cur := head
		for cur != end {
			if cur.Value < value && cur.Next.Value >= value {
				next := cur.Next
				cur.Next = &Link{Value: value}
				cur.Next.Next = next
				break
			}
			cur = cur.Next
		}
	}

	if value <= head.Value || value >= end.Value {
		end.Next = &Link{Value: value}
		end.Next.Next = head
	}

	return head
}
