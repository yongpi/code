package tree

func FindSuccessorNode(node *Tree) *Tree {
	if node.Right != nil {
		sn := node.Right
		for sn.Left != nil {
			sn = sn.Left
		}

		return sn
	}

	parent := node.Parent
	if parent == nil {
		return nil
	}

	if parent.Left == node {
		return parent
	}

	son := parent
	parent = parent.Parent
	for parent != nil {
		if parent.Left == son {
			return parent
		}

		son = parent
		parent = parent.Parent
	}

	return nil
}

func FindSuccessorNode2(node *Tree) *Tree {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return node.Right
	}

	parent := node.Parent
	if parent == nil {
		return nil
	}

	if parent.Left == node {
		return parent
	}

	for parent.Right == node {
		grand := parent.Parent
		if grand == nil {
			return nil
		}

		if grand.Left == parent {
			return grand
		}

		if grand.Right == parent {
			node = parent
			parent = grand
			continue
		}

		return nil
	}

	return nil
}
