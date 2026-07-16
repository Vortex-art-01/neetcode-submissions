func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))

	fmt.Println("temperatures", temperatures)

	for i := 0; i < len(res); i++ {
		stack := get_stack(temperatures, i)

		fmt.Println("for item", temperatures[i])
		fmt.Println("stack", stack)

		day_count := 0
		count := 0
		for !stack.Empty() {
			count++
			if temperatures[i] < stack.Pop() {
				day_count = count
				break
			} 
		} 
		res[i] = day_count
	}

	return res
}

func get_stack(temperatures []int, cur_idx int) stack {
	res := stack{
		arr: []int{},
	}
	// fmt.Println("temperatures int get_stack", temperatures)
	// fmt.Println("cur_idx", cur_idx)

	for i := len(temperatures)-1; i > cur_idx; i-- {
		// fmt.Println("for i", i)
		// fmt.Println("temperatures[i]", temperatures[i])

		res.Push(temperatures[i])
	}

	return res
}

type stack struct {
	arr []int
}

func (s *stack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stack) Pop() int {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

func (s stack) Empty() bool {
	return len(s.arr) == 0
}

