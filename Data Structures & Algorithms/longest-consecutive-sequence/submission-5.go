func longestConsecutive(nums []int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}

	max := 0;

	for _,num := range nums {
		// fmt.Println("num", num)
		// fmt.Println("num-1", num-1)

		if _, found := hash[num-1]; !found {
			plus := 0

			for {
				if _, found := hash[num+plus]; !found {
				// fmt.Println("NOT FOUND num", num)
					break
				}
				// fmt.Println("num", num)
				// fmt.Println("hash", hash)
				// fmt.Println("found", num+plus)
				plus++
			}

			if plus > max {
				max = plus
			}
		} 
	}

	return max
}
