package link

func CopyRandom(head *Link) *Link {
	if head == nil {
		return nil
	}

	hm := make(map[*Link]*Link)

	tmp := *head
	ch := &tmp
	ch.Value = ch.Value * 10
	hm[head] = ch

	node := head
	for node.Next != nil {
		node = node.Next
		tmp := *node
		tmp.Value = tmp.Value * 10
		hm[node] = &tmp
		ch.Next = hm[node]
		ch = ch.Next
	}

	node = head
	for node != nil {
		hm[node].Random = hm[node.Random]
		node = node.Next
	}

	return hm[head]
}

func CopyRandom2(head *Link) *Link {
	if head == nil {
		return head
	}

	node := head
	for node != nil {
		next := node.Next

		cp := *node
		cp.Value = cp.Value * 10
		node.Next = &cp
		node.Next.Next = next

		node = next
	}

	nh := head.Next
	node = nh
	for node != nil {
		if node.Random != nil {
			node.Random = node.Random.Next
		}
		if node.Next != nil {
			node.Next = node.Next.Next
		}

		node = node.Next
	}

	return nh
}
