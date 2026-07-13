import (
	"slices"
)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	noDbl := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		noDbl[nums[i]] = true
	}

	res := []int{}
	for num := range noDbl {
		_, found_less := noDbl[num-1];
		_, found_greater := noDbl[num+1];

		if !found_less && !found_greater {
			delete(noDbl, num)
		} else {
			res = append(res, num)
		}
	}
	fmt.Println("after clear",noDbl)

	slices.Sort(res)

	fmt.Println("after sort",res)
	curCount := 1;
	maxCount := 1
	for i := 1; i < len(res); i++ {
		if res[i-1] == res[i]-1 {
			curCount++
		} else {
			if curCount > maxCount {
				maxCount = curCount
			}
			curCount = 1
		}
	}
	
	if curCount > maxCount {
		maxCount = curCount
	}

	fmt.Println("curCount",curCount)
	fmt.Println("maxCount",maxCount)

	return maxCount
}
