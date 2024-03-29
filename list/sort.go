package list

func MiniNotSortSubList(nums []int) []int {
	minValue := nums[len(nums)-1]
	begin := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > minValue {
			begin = i
			continue
		}
		minValue = nums[i]
	}

	if begin == len(nums)-1 {
		return []int{}
	}

	end := 0
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < maxValue {
			end = i
			continue
		}
		maxValue = nums[i]
	}

	if end == 0 {
		return []int{}
	}

	return nums[begin : end+1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func MaxCanMergeSub(nums []int) []int {
	var begin, end, count int
	for i := 0; i < len(nums); i++ {
		maxValue, minValue := nums[i], nums[i]
		hm := make(map[int]struct{})
		for j := i; j < len(nums); j++ {
			if _, ok := hm[nums[j]]; ok {
				break
			}

			hm[nums[j]] = struct{}{}
			maxValue = Max(maxValue, nums[j])
			minValue = Min(minValue, nums[j])

			if (j-i == maxValue-minValue) && (j-i+1) > count {
				begin = i
				end = j
				count = j - i + 1
			}
		}
	}

	return nums[begin : end+1]
}

func SortN(nums []int) []int {
	n := len(nums)
	for i := 1; i <= n; i++ {
		for nums[i-1] != i {
			value := nums[i-1]
			nums[value-1], nums[i-1] = value, nums[value-1]
		}
	}

	return nums
}

func SortOddEven(nums []int) []int {
	odd, even := 1, 0
	n := len(nums) - 1

	for odd <= n && even <= n {
		if nums[n]%2 == 0 {
			nums[n], nums[even] = nums[even], nums[n]
			even += 2
		} else {
			nums[n], nums[odd] = nums[odd], nums[n]
			odd += 2
		}
	}

	return nums
}

func SortZeroOneTwo(nums []int) []int {
	zero, two := 0, len(nums)-1
	index := 0

	for index < two {
		if nums[index] == 0 {
			nums[index], nums[zero] = nums[zero], nums[index]
			zero++
		} else if nums[index] == 2 {
			nums[index], nums[two] = nums[two], nums[index]
			two--
		} else {
			index++
		}
	}

	return nums
}
