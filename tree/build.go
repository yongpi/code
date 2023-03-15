package tree

func BuildPost(nums []int) *Tree {
	if len(nums) == 0 {
		return nil
	}

	root := &Tree{Value: nums[len(nums)-1]}
	mid := 0
	for mid < len(nums)-1 {
		if nums[mid] > nums[len(nums)-1] {
			break
		}

		mid++
	}

	root.Left = BuildPost(nums[:mid])
	root.Right = BuildPost(nums[mid : len(nums)-1])

	return root
}

func BuildSearchMid(nums []int) *Tree {
	if len(nums) == 0 {
		return nil
	}

	index := len(nums) / 2

	root := &Tree{Value: nums[index]}
	root.Left = BuildSearchMid(nums[:index])
	root.Right = BuildSearchMid(nums[index+1:])

	return root
}

func BuildPostList(pre, mid []int) []int {
	post := make([]int, len(pre))
	stack := make([]int, 0)
	for key, _ := range pre {
		stack = append(stack, key)
	}

	makePostList(pre, mid, post, stack)

	return post
}

func makePostList(pre, mid, post, stack []int) []int {
	if len(pre) == 0 || len(mid) == 0 || len(stack) == 0 {
		return stack
	}

	root := pre[0]
	i := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	post[i] = root

	var ml int
	for _, value := range mid {
		if value != root {
			ml++
		}
		if value == root {
			break
		}
	}

	stack = makePostList(pre[1+ml:], mid[ml+1:], post, stack)
	stack = makePostList(pre[1:ml+1], mid[:ml], post, stack)

	return stack
}
