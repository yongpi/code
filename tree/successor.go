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
