package str

type Node struct {
	Value byte
	Child map[byte]*Node
	Count int
}

func Insert(root *Node, data string) {
	if root.Child == nil {
		root.Child = make(map[byte]*Node)
	}

	child := root.Child
	for i := 0; i < len(data); i++ {
		if item, ok := child[data[i]]; ok {
			item.Count++
			child = item.Child
			continue
		}

		item := &Node{
			Value: data[i],
			Child: make(map[byte]*Node),
			Count: 1,
		}

		child[data[i]] = item
		child = item.Child
	}
}

func Exist(root *Node, data string) bool {
	if root == nil {
		return false
	}

	child := root.Child
	for i := 0; i < len(data); i++ {
		item, ok := child[data[i]]
		if !ok {
			return false
		}
		child = item.Child
	}

	return true
}

func Delete(root *Node, data string) {
	if root == nil {
		return
	}

	if !Exist(root, data) {
		return
	}

	child := root.Child
	for i := 0; i < len(data); i++ {
		item, ok := child[data[i]]
		if !ok {
			return
		}

		if item.Count == 1 {
			delete(child, data[i])
			return
		}

		item.Count--
		child = item.Child
	}
}

func PreCount(root *Node, pre string) int {
	if root == nil {
		return 0
	}

	child := root.Child
	var count int
	for i := 0; i < len(pre); i++ {
		item, ok := child[pre[i]]
		if !ok {
			return 0
		}

		child = item.Child
		count = item.Count
	}

	return count
}
