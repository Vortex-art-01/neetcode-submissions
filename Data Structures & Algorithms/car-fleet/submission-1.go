func carFleet(target int, position []int, speed []int) int {
	pairs := [][2]int{}

	for i := 0; i < len(position); i++ {
		pairs = append(pairs, [2]int{position[i], speed[i]})
	}

	sort.Slice(pairs, func(i,j int) bool {
		return pairs[i][0] > pairs[j][0]
	}) 

// fmt.Println("pairs sorted", pairs)
	stack := []float64{}

	for _, pair := range pairs {
		time := float64(target - pair[0]) / float64(pair[1])
		if len(stack) == 0 {
			stack = append(stack, time)
		} else if time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}
// fmt.Println("stack", stack)

	return len(stack)
}
