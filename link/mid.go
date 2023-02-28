package link

func DeleteMid(head *Link) *Link {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head.Next
	}

	pre := head
	node := head.Next.Next

	for node.Next != nil && node.Next.Next != nil {
		node = node.Next.Next
		pre = pre.Next
	}

	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}

	return head
}
