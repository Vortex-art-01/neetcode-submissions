func topKFrequent(nums []int, k int) []int {
	freqs := map[int]int{}
	for _, val := range nums {
		freqs[val]++
	}

	res := []int{}
	
    for i := 0; i < k; i++ {
        first := true
        max := []int{}

        for num, freq := range freqs {
            if first || freq > max[1] {
                max = []int{num, freq}
            }
            first = false
        }

        if len(max) != 0 {
            delete(freqs, max[0])
        }

        res = append(res, max[0])
    }

	return res
}
