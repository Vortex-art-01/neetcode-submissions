func topKFrequent(nums []int, k int) []int {
	freqs := map[int]int{}
    count := make([][]int, len(nums)+1, len(nums)+1)
	for _, val := range nums {
		freqs[val]++
	}
    fmt.Println(freqs)

    for num, freq := range freqs {
        count[freq] = append(count[freq], num)
    }
    fmt.Println(count)
    res := []int{}

    for i := len(count)-1; i >= 0; i-- {
        if len(count[i]) != 0 {
            in_arr := count[i]
            
            for j := 0; j < len(in_arr); j++ {
                res = append(res, in_arr[j])

                if (len(res) == k) {
                    return res
                }
            }
        }
    }

    return res


    return res
}
