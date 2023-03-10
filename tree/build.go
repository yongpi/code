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
