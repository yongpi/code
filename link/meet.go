package link

func TwoMeet(l1, l2 *Link) (bool, int) {
	var len1, end1, len2, end2 int

	node := l1
	for node != nil {
		len1++
		end1 = node.Value
		node = node.Next
	}

	node = l2
	for node != nil {
		len2++
		end2 = node.Value
		node = node.Next
	}

	if end1 != end2 {
		return false, 0
	}

	for l2 != nil && len2 > len1 {
		l2 = l2.Next
		len2--
	}

	for l1 != nil && len1 > len2 {
		l1 = l1.Next
		len1--
	}

	for l1 != nil && l2 != nil {
		if l1.Value == l2.Value {
			return true, l1.Value
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return false, 0
}

func CircleMeet(head *Link) int {
	slow := head.Next
	fast := head.Next.Next
	for slow != fast && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = head
	for slow != fast && slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Value
}
