package link

func Add(l1, l2 *Link) *Link {
	if l1 == nil && l2 == nil {
		return nil
	}

	var pre *Link
	for l1 != nil {
		next := l1.Next
		l1.Next = pre
		pre = l1
		l1 = next
	}

	var pre2 *Link
	for l2 != nil {
		next := l2.Next
		l2.Next = pre2
		pre2 = l2
		l2 = next
	}

	var nh, node *Link
	var count int

	for pre != nil || pre2 != nil {
		value := count
		if pre != nil {
			value += pre.Value
			pre = pre.Next
		}
		if pre2 != nil {
			value += pre2.Value
			pre2 = pre2.Next
		}

		count = value / 10
		value = value % 10

		if nh == nil {
			nh = &Link{Value: value}
			node = nh
		} else {
			node.Next = &Link{Value: value}
			node = node.Next
		}
	}

	if count > 0 {
		node.Next = &Link{Value: count}
	}

	var head *Link
	for nh != nil {
		next := nh.Next
		nh.Next = head
		head = nh
		nh = next
	}

	return head
}
