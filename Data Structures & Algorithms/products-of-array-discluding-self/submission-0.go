func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums), len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = getSumExcludeOneIdx(nums, i)
	}
	return res
}

func getSumExcludeOneIdx(nums []int, idx int) int {
	res := 1;
	
	for i := 0; i < len(nums); i++ {
		if i == idx {
			continue
		}
		res = res * nums[i]
	}

	return res
}
