func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums), len(nums))
	suffix := make([]int, len(nums), len(nums))

	prefix[0], suffix[len(suffix)-1] = 1, 1

	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(suffix)-2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	res := make([]int, len(nums), len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = prefix[i] * suffix[i]
	}
	
	return res
}

