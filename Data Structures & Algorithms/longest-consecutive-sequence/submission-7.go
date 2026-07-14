func longestConsecutive(nums []int) int {
	hash := map[int]bool{}
	
	for _, num := range nums {
		hash[num] = true
	}

	longest := 0

	for _, num := range nums {
		if _, found := hash[num-1]; !found {
			count := 0

			for {
				if _, found := hash[num+count]; !found {
					if count > longest {
						longest = count
					}

					break
				}
				count++
			}
		} 
	}
	return longest
}

