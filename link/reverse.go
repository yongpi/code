package link

func Reverse(head *Link) *Link {
	var pre *Link
	node := head

	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre
}

func ReverseV2(head *Link) *Link {
	var pre *Link
	node := head

	for node != nil {
		next := node.Next
		node.Next = pre
		node.Before = next
		pre = node
		node = next
	}

	return pre
}

func ReverseFromTo(head *Link, from, to int) *Link {
	if from >= to || head == nil {
		return head
	}

	var fl, tl *Link
	node := head

	for node != nil {
		to--
		from--
		if from == 1 {
			fl = node
		}

		if to == 0 {
			tl = node
			break
		}

		node = node.Next
	}

	if to > 0 {
		return head
	}

	node = head
	if fl != nil {
		node = fl.Next
	}

	tln := tl.Next
	tl.Next = nil

	var pre *Link
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	if fl != nil {
		fl.Next = pre
	}

	node.Next = tln

	if fl == nil {
		return pre
	}
	return head
}
