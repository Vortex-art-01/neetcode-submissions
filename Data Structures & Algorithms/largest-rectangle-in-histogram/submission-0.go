func largestRectangleArea(heights []int) int {
	m := map[int]bool{}

	for _, num := range heights {
		m[num] = true
	}

	maxRect := 0

	for num, _ := range m {
		width := 0

		for i := 0; i < len(heights); i++ {
			if heights[i] >= num {
				width++
			} else {
				localRect := width * num
				if localRect > maxRect {
					maxRect = localRect
				}

				width = 0
			}
		}

		// fmt.Println("num",num,"width", width)
		localRect := width * num
		if localRect > maxRect {
			maxRect = localRect
		}
	} 

	return maxRect
}
