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

func IsPostArray(list []int) bool {
	if len(list) == 0 {
		return false
	}

	var mid int
	root := list[len(list)-1]
	for i := 0; i < len(list)-1; i++ {
		if list[i] < root {
			mid = i
		} else {
			break
		}
	}

	if mid+1 >= len(list)-1 {
		return true
	}

	for i := mid + 1; i < len(list)-1; i++ {
		if list[i] <= root {
			return false
		}
	}

	return IsPostArray(list[:mid+1]) && IsPostArray(list[mid+1:len(list)-1])
}

func BuildPost2(nums []int) *Tree {
	if len(nums) == 0 {
		return nil
	}

	root := &Tree{Value: nums[len(nums)-1]}
	if len(nums) == 1 {
		return root
	}

	var mid int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < root.Value {
			mid = i
		} else {
			break
		}
	}

	root.Left = BuildPost2(nums[:mid+1])
	root.Right = BuildPost2(nums[mid+1 : len(nums)-1])

	return root
}

func BuildPostList2(pre, mid []int) []int {
	post := make([]int, len(pre))
	index := len(pre) - 1

	var makePost func(pre, mid, post []int, ci int) int
	makePost = func(pre, mid, post []int, ci int) int {
		if len(pre) == 0 {
			return ci
		}

		root := pre[0]
		post[ci] = root
		ci--

		if len(pre) == 1 {
			return ci
		}

		var leftLen int
		for i := 0; i < len(mid); i++ {
			if mid[i] == root {
				leftLen = i
				break
			}
		}

		ci = makePost(pre[1+leftLen:], mid[leftLen+1:], post, ci)
		ci = makePost(pre[1:leftLen+1], mid[:leftLen], post, ci)

		return ci
	}

	makePost(pre, mid, post, index)
	return post
}
