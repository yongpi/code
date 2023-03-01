package link

func IsHw(head *Link) bool {
	if head == nil || head.Next == nil {
		return true
	}

	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Value)
		head = head.Next
	}

	max := len(stack) - 1
	mid := len(stack) / 2
	for i := 0; i < mid; i++ {
		if stack[i] != stack[max-i] {
			return false
		}
	}

	return true
}

func IsHw2(head *Link) (bool, *Link) {
	if head == nil || head.Next == nil {
		return true, head
	}

	mid := head
	cur := head

	for cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next.Next
		mid = mid.Next
	}

	head2 := mid.Next
	mid.Next = nil

	var pre *Link
	for head2 != nil {
		next := head2.Next
		head2.Next = pre
		pre = head2
		head2 = next
	}

	head2 = pre
	cur = head

	defer func() {
		pre = nil
		for head2 != nil {
			next := head2.Next
			head2.Next = pre
			pre = head2
			head2 = next
		}
		head2 = pre
		mid.Next = head2
	}()

	for pre != nil {
		if cur == nil || pre.Value != cur.Value {
			return false, head
		}

		pre = pre.Next
		cur = cur.Next
	}

	return true, head
}
