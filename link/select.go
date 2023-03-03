package link

func SelectionSort(head *Link) *Link {
	if head == nil || head.Next == nil {
		return head
	}

	nh := head
	node := head
	pre := head
	for node.Next != nil {
		if node.Next.Value < nh.Value {
			pre = node
			nh = node.Next
		}
		node = node.Next
	}

	if pre != nh && pre.Next != nil {
		pre.Next = pre.Next.Next
	}

	oh := head
	if nh == head {
		oh = nh.Next
		nh.Next = nil
	}

	nn := nh
	for oh.Next != nil {
		node = oh
		cur, cp := oh, oh

		for node.Next != nil {
			if node.Next.Value < cur.Value {
				cp = node
				cur = node.Next
			}

			node = node.Next
		}

		if cur != oh && cp.Next != nil {
			cp.Next = cp.Next.Next
		}

		if cur == oh {
			oh = cur.Next
			cur.Next = nil
		}

		nn.Next = cur
		nn = cur
	}

	nn.Next = oh

	return nh

}
