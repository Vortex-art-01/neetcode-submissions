func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for idx, val := range nums {
		want := target - val 

		if found_idx, found := m[want]; found {
			return []int{found_idx, idx}
		}
		m[val] = idx
	}
	return nil
}
