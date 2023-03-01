package link

func MakeThree(head *Link, num int) *Link {
	if head == nil || head.Next == nil {
		return head
	}

	var one, two, three *Link
	for head != nil {
		next := head.Next
		value := head.Value
		if value < num {
			if one == nil {
				one = head
			} else {
				one.Next = head
			}
		}

		if value == num {
			if two == nil {
				two = head
			} else {
				two.Next = head
			}
		}

		if value > num {
			if three == nil {
				three = head
			} else {
				three.Next = head
			}
		}

		head.Next = nil
		head = next
	}

	list := []*Link{one, two, three}
	var node, nh *Link
	for _, item := range list {
		if node == nil && item != nil {
			nh = item
			node = item
			for node.Next != nil {
				node = node.Next
			}
		} else if node != nil && item != nil {
			node.Next = item
			for node.Next != nil {
				node = node.Next
			}
		}

	}

	return nh
}
