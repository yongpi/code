package link

func DeleteLastN(head *Link, n int) *Link {
	node := head

	for node != nil && n > 0 {
		node = node.Next
		n--
	}

	if n > 0 {
		panic("n too max")
	}

	if node == nil {
		head = head.Next
		return head
	}

	pre := head
	for node != nil && node.Next != nil {
		node = node.Next
		pre = pre.Next
	}

	if pre != nil && pre.Next != nil {
		pre.Next = pre.Next.Next
	}

	return head
}

func DeleteLastN2(head *Link, n int) *Link {
	node := head

	for node != nil && n > 0 {
		n--
		node = node.Next
	}

	if n > 0 {
		panic("n too max")
	}

	if node == nil {
		next := head.Next
		next.Before = nil
		return next
	}

	pre := head
	for node != nil && node.Next != nil {
		node = node.Next
		pre = pre.Next
	}

	next := pre.Next
	if next != nil {
		pre.Next = next.Next
		if next.Next != nil {
			next.Next.Before = pre
		}
	}

	return head
}
