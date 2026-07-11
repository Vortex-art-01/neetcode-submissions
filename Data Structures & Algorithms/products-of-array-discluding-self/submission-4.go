func productExceptSelf(nums []int) []int {
	countZero := 0
	sum := 1

	for _, num := range nums {
		if num == 0 {
			countZero++
		} else {
			sum *= num
		}
	}

	res := make([]int, len(nums), len(nums))

	if countZero > 1 {
		return res
	}

	for i := 0; i < len(res); i++ {
		if countZero == 1 {
			if nums[i] != 0 {
				res[i] = 0
			} else {
				res[i] = sum
			}
		} else {
			res[i] = sum/nums[i]
		}
	}
	return res
}

